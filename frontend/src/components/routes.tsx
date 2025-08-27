/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import {BeakerIcon, CollectionIcon, CubeTransparentIcon, FilterIcon, HomeIcon, ScaleIcon,} from '@heroicons/react/outline';
import type {NavLinkProps} from '@redpanda-data/ui/dist/components/Nav/NavLink';
import React, {Fragment, type FunctionComponent} from 'react';
import {Redirect, Route} from 'react-router';
import {Switch} from 'react-router-dom';
import {api} from '../state/backendApi';
import type {UserPermissions} from '../state/restInterfaces';
import {Feature, type FeatureEntry, isSupported, shouldHideIfNotSupported} from '../state/supportedFeatures';
import {uiState} from '../state/uiState';
import {AnimatePresence} from '../utils/animationProps';
import {type AppFeature, AppFeatures} from '../utils/env';
import {Section} from './misc/common';
import type {PageComponentType, PageProps} from './pages/Page';
import GroupDetails from './pages/consumers/Group.Details';
import GroupList from './pages/consumers/Group.List';
import {BrokerDetails} from './pages/overview/Broker.Details';
import Overview from './pages/overview/Overview';
import QuotasList from './pages/quotas/Quotas.List';
import ReassignPartitions from './pages/reassign-partitions/ReassignPartitions';
import EditSchemaCompatibilityPage from './pages/schemas/EditCompatibility';
import {SchemaAddVersionPage, SchemaCreatePage} from './pages/schemas/Schema.Create';
import SchemaDetailsView from './pages/schemas/Schema.Details';
import SchemaList from './pages/schemas/Schema.List';
import TopicDetails from './pages/topics/Topic.Details';
import TopicList from './pages/topics/Topic.List';
import {TopicProducePage} from './pages/topics/Topic.Produce';

//
//	Route Types
//
type IRouteEntry = PageDefinition<any>;

export interface PageDefinition<TRouteParams = {}> {
  title: string;
  path: string;
  pageType: PageComponentType<TRouteParams> | FunctionComponent<TRouteParams>;
  routeJsx: JSX.Element;
  icon?: (props: React.ComponentProps<'svg'>) => JSX.Element;
  menuItemKey?: string; // set by 'CreateRouteMenuItems'
  visibilityCheck?: () => MenuItemState;
}

// Generate content for <Menu> from all routes
export function createVisibleSidebarItems(entries: IRouteEntry[]): NavLinkProps[] {
  return entries
    .map((entry) => {
      // Menu entry for Page
      if (entry.path.includes(':')) return null; // only root-routes (no param) can be in menu
      if (!entry.icon) return null; // items without icon do not appear in the sidebar

      let isEnabled = true;
      let disabledText: JSX.Element = <Fragment key={entry.title}/>;
      if (entry.visibilityCheck) {
        const visibility = entry.visibilityCheck();
        if (!visibility.visible) return null;

        isEnabled = visibility.disabledReasons.length === 0;
        if (!isEnabled) disabledText = disabledReasonText[visibility.disabledReasons[0]];
      }
      const isDisabled = !isEnabled;

      return {
        title: entry.title as string,
        to: entry.path as string,
        icon: entry.icon as any,
        isDisabled: isDisabled as boolean,
        disabledText: disabledText as unknown as string,
      };
    })
    .filter((x) => x != null && x !== undefined) as NavLinkProps[];
}

// Convert routes to <Route/> JSX declarations
function EmitRouteViews(entries: IRouteEntry[]): JSX.Element[] {
  return entries.map((e) => e.routeJsx);
}

export const RouteView = () => (
  <AnimatePresence mode="wait">
    <Switch>
      {/* Index */}
      <Route exact path="/" render={() => <Redirect to="/overview"/>}/>

      {/* Emit all <Route/> elements */}
      {EmitRouteViews(APP_ROUTES)}

      <Route
        render={(rp) => {
          uiState.pageTitle = '404';
          return (
            <Section title="404">
              <div>
                <h4>Path:</h4> <span>{rp.location.pathname}</span>
              </div>
              <div>
                <h4>Query:</h4>
                <pre>{JSON.stringify(rp.location.search, null, 4)}</pre>
              </div>
            </Section>
          );
        }}
      />
    </Switch>
  </AnimatePresence>
);

enum DisabledReasons {
  notSupported = 0, // kafka cluster version too low
  noPermission = 1, // user doesn't have permissions to use the feature
  enterpriseFeature = 2,
}

const disabledReasonText: { [key in DisabledReasons]: JSX.Element } = {
  [DisabledReasons.noPermission]: (
    <span>
      You don't have premissions
      <br/>
      to view this page.
    </span>
  ),
  [DisabledReasons.notSupported]: (
    <span>
      The Kafka cluster does not
      <br/>
      support this feature.
    </span>
  ),
  [DisabledReasons.enterpriseFeature]: <span>This feature requires an enterprise license.</span>,
} as const;

interface MenuItemState {
  visible: boolean;
  disabledReasons: DisabledReasons[];
}

function MakeRoute<TRouteParams>(
  path: string,
  page: PageComponentType<TRouteParams> | FunctionComponent<TRouteParams>,
  title: string,
  icon?: (props: React.ComponentProps<'svg'>) => JSX.Element,
  exact = true,
  showCallback?: () => MenuItemState,
): PageDefinition<TRouteParams> {
  const route: PageDefinition<TRouteParams> = {
    title,
    path,
    pageType: page,
    routeJsx: null as unknown as JSX.Element, // will be set below
    icon,
    visibilityCheck: showCallback,
  };

  // todo: verify that path and route params match
  route.routeJsx = (
    <Route
      path={route.path}
      key={route.title}
      exact={exact ? true : undefined}
      render={(rp) => {
        const matchedPath = rp.match.url;
        const {...params} = rp.match.params;

        if (uiState.currentRoute && uiState.currentRoute.path !== route.path) {
          //kconsole.log('switching route: ' + routeStr(ui.currentRoute) + " -> " + routeStr(route));
        }

        const pageProps: PageProps<TRouteParams> = {
          matchedPath,
          ...params,
        } as PageProps<TRouteParams>;

        uiState.currentRoute = route;
        return <route.pageType {...pageProps} />;
      }}
    />
  );

  return route;
}

function routeVisibility(
  visible: boolean | (() => boolean),
  requiredFeatures?: FeatureEntry[],
  requiredPermissions?: UserPermissions[],
  requiredAppFeatures?: AppFeature[],
): () => MenuItemState {
  return () => {
    let v = typeof visible === 'boolean' ? visible : visible();

    const disabledReasons: DisabledReasons[] = [];
    if (requiredFeatures)
      for (const f of requiredFeatures) {
        if (!isSupported(f)) {
          if (shouldHideIfNotSupported(f)) {
            v = false;
          } else {
            disabledReasons.push(DisabledReasons.notSupported);
          }
          break;
        }
      }

    if (requiredPermissions && api.userData)
      for (const p of requiredPermissions) {
        const hasPermission = api.userData[p];
        if (!hasPermission) {
          disabledReasons.push(DisabledReasons.noPermission);
          break;
        }
      }

    if (requiredAppFeatures) {
      for (const f of requiredAppFeatures)
        if (AppFeatures[f] === false) {
          disabledReasons.push(DisabledReasons.enterpriseFeature);
          break;
        }
    }

    return {
      visible: v,
      disabledReasons: disabledReasons,
    };
  };
}

//
// Route Definitions
// If a route has one or more parameters it will not be shown in the main menu (obviously, since the parameter would have to be known!)
//
export const APP_ROUTES: IRouteEntry[] = [
  MakeRoute<{}>('/overview', Overview, 'Overview', HomeIcon),
  MakeRoute<{ brokerId: string }>('/overview/:brokerId', BrokerDetails, 'Broker Details'),

  MakeRoute<{}>('/topics', TopicList, 'Topics', CollectionIcon),
  MakeRoute<{ topicName: string }>('/topics/:topicName', TopicDetails, 'Topics'),
  MakeRoute<{ topicName: string }>('/topics/:topicName/produce-record', TopicProducePage, 'Produce Record'),

  MakeRoute<{}>('/schema-registry', SchemaList, 'Schema Registry', CubeTransparentIcon),
  MakeRoute<{}>('/schema-registry/create', SchemaCreatePage, 'Create schema'),
  MakeRoute<{ subjectName: string }>(
    '/schema-registry/subjects/:subjectName/add-version',
    SchemaAddVersionPage,
    'Add version',
  ),
  MakeRoute<{ subjectName: string }>('/schema-registry/subjects/:subjectName', SchemaDetailsView, 'Schema Registry'),
  MakeRoute<{ subjectName: string }>(
    '/schema-registry/edit-compatibility',
    EditSchemaCompatibilityPage,
    'Edit Schema Compatibility',
  ),
  MakeRoute<{ subjectName: string }>(
    '/schema-registry/subjects/:subjectName/edit-compatibility',
    EditSchemaCompatibilityPage,
    'Edit Schema Compatibility',
  ),

  MakeRoute<{}>(
    '/groups',
    GroupList,
    'Consumer Groups',
    FilterIcon,
    undefined,
    routeVisibility(true, [Feature.ConsumerGroups]),
  ),
  MakeRoute<{ groupId: string }>('/groups/:groupId/', GroupDetails, 'Consumer Groups'),

  MakeRoute<{}>(
    '/quotas',
    QuotasList,
    'Quotas',
    ScaleIcon,
    true,
    routeVisibility(true, [Feature.GetQuotas], ['canListQuotas']),
  ),

  MakeRoute<{}>(
    '/reassign-partitions',
    ReassignPartitions,
    'Reassign Partitions',
    BeakerIcon,
    false,
    routeVisibility(
      true,
      [Feature.GetReassignments, Feature.PatchReassignments],
      ['canPatchConfigs', 'canReassignPartitions'],
      ['REASSIGN_PARTITIONS'],
    ),
  ),
].filterNull();
