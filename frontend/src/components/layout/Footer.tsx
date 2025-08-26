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

import {FaGithub} from 'react-icons/fa';
import env, {getBuildDate, IsCI, IsDev} from '../../utils/env';

export const VersionInfo = () => {
  const appName = 'Kafka Console';
  let mode = '';
  if (IsDev) mode = ' - DEV';
  if (IsCI) mode = ' - CI';

  if (env.REACT_APP_CONSOLE_PLATFORM_VERSION) mode += ` (Platform Version ${env.REACT_APP_CONSOLE_PLATFORM_VERSION})`;

  let ref = env.REACT_APP_CONSOLE_GIT_REF;
  if (!ref || ref === 'master') ref = '';

  const sha = IsDev ? '<no git sha in dev>' : env.REACT_APP_CONSOLE_GIT_SHA.slice(0, 7);

  const buildDate = IsDev ? new Date() : getBuildDate();

  return (
    <>
      <div data-testid="versionTitle" className="versionTitle">
        {appName} {mode}
      </div>
      <div className="versionDate">
        (built {buildDate?.toLocaleDateString('en-US', {month: 'long', day: '2-digit', year: 'numeric'})})
      </div>
      <div className="versionGitData">
        {ref} {sha}
      </div>
    </>
  );
};

export const AppFooter = () => {
  const gitHub = (link: string, title: string) => (
    <>
      <a href={link} title={title} target="_blank" rel="noopener noreferrer">
        <FaGithub/>
      </a>
    </>
  );

  return (
    <footer className="footer">
      {/* Social Media Links */}
      <div className="links">
        {
          gitHub('https://github.com/xxxcrel/kafka-console', "Visit Kafka Console's GitHub repository")
        }
      </div>
      {/* Version Info */}
      <div className="versionText">
        <VersionInfo/>
      </div>
    </footer>
  );
};

export default AppFooter;
