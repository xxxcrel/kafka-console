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

import { ThreeBarsIcon, XIcon } from '@primer/octicons-react';
import { Button, Input, Tooltip } from '@redpanda-data/ui';
import { arrayMoveMutable } from 'array-move';
import { type IReactionDisposer, autorun, computed, makeObservable, observable } from 'mobx';
import { observer } from 'mobx-react';
import { Component, useState } from 'react';
import { DragDropContext, Draggable, type DropResult, Droppable, type ResponderProvided } from 'react-beautiful-dnd';

@observer
export class CommaSeparatedStringList extends Component<{
  defaultValue: string;
  // renderItem: (item: string, index: number, ar: string[]) => JSX.Element,
  onChange: (list: string) => void;

  locale?: {
    addInputPlaceholder?: string;
    addButtonText?: string;
  };
}> {
  @observable data: { id: string }[];
  @observable newEntry: string | null = null;
  @observable newEntryError: string | null = null;

  reactionDisposer: IReactionDisposer | undefined;

  constructor(p: any) {
    super(p);
    makeObservable(this);

    if (!this.data)
      this.data = this.props.defaultValue ? this.props.defaultValue.split(',').map((x) => ({ id: x.trim() })) : [];

    this.reactionDisposer = autorun(() => {
      const list = this.commaSeperatedList;
      this.props.onChange(list);
    });
  }

  componentWillUnmount() {
    this.reactionDisposer?.();
  }

  @computed get commaSeperatedList(): string {
    const str = this.data.map((x) => x.id).join(',');
    return str;
  }

  render() {
    return (
      <div className="stringList" style={{ maxWidth: '500px' }}>
        <this.AddButton />
        <List observableAr={this.data} renderItem={(item, index) => <this.Item item={item} index={index} />} />
      </div>
    );
  }

  Item = observer((props: { item: { id: string }; index: number }): JSX.Element => {
    const { item, index } = props;

    const [hasFocus, setHasFocus] = useState(false);
    const [valuePending, setValuePending] = useState('');

    return (
      <>
        {/* Input */}
        <Tooltip label="[Enter] confirm, [ESC] cancel" isOpen={hasFocus} placement="top" hasArrow={true}>
          <Input
            className="ghostInput"
            size="sm"
            style={{ flexGrow: 1, flexBasis: '400px' }}
            onFocus={() => {
              setValuePending(item.id);
              setHasFocus(true);
            }}
            onBlur={() => {
              setHasFocus(false);
            }}
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                // can we rename that entry?
                if (this.data.any((x) => x.id === valuePending)) {
                  // no, already exists
                  e.stopPropagation();
                  return;
                }

                if (/^[a-z][a-z_\d]*$/i.test(valuePending) === false) {
                  // no, invalid characters
                  e.stopPropagation();
                  return;
                }

                item.id = valuePending;
                (e.target as HTMLElement).blur();
              } else if (e.key === 'Escape') {
                (e.target as HTMLElement).blur();
              }
            }}
            value={hasFocus ? valuePending : item.id}
            onChange={(e) => setValuePending(e.target.value)}
            spellCheck={false}
          />
        </Tooltip>

        {/* Delete */}
        <span
          className="deleteButton"
          onClick={() => {
            this.data.splice(index, 1);
          }}
        >
          <XIcon />
        </span>
      </>
    );
  });

  AddButton = observer(() => {
    return (
      <div className="createEntryRow">
        <div className={`inputWrapper${this.newEntryError ? ' hasError' : ''}`} style={{ height: '100%' }}>
          <Input
            style={{ flexGrow: 1, height: '100%', flexBasis: '260px' }}
            value={this.newEntry ?? ''}
            onChange={(e) => {
              this.newEntry = e.target.value;

              this.newEntryError = null;

              if (!this.newEntry) return;

              if (this.data.any((x) => x.id === this.newEntry)) this.newEntryError = 'Entry already exists';

              if (/^[a-z][a-z_\d]*$/i.test(this.newEntry) === false)
                this.newEntryError = 'Name is not valid (only letters, digits, underscore)';
            }}
            placeholder={this.props.locale?.addInputPlaceholder ?? 'Enter a name...'}
            spellCheck={false}
          />

          <div className="validationFeedback">{this.newEntryError ?? null}</div>
        </div>

        <Button
          style={{ padding: '0px 16px', height: '100%', minWidth: '120px' }}
          variant="solid"
          disabled={this.newEntryError != null || !this.newEntry || this.newEntry.trim().length === 0}
          onClick={() => {
            // biome-ignore lint/style/noNonNullAssertion: not touching MobX observables
            this.data.push({ id: this.newEntry! });
            this.newEntry = null;
          }}
        >
          Add
        </Button>
      </div>
    );
  });
}

@observer
export class List<T extends { id: string }> extends Component<{
  observableAr: T[];
  renderItem: (item: T, index: number) => JSX.Element;
}> {
  render() {
    const { observableAr: list, renderItem } = this.props;

    const onDragEnd = (result: DropResult, _provided: ResponderProvided) => {
      if (!result.destination) return;
      arrayMoveMutable(this.props.observableAr, result.source.index, result.destination.index);
    };

    return (
      <div className="reorderableList">
        <DragDropContext onDragEnd={onDragEnd}>
          <Droppable droppableId="droppable" ignoreContainerClipping={false}>
            {(droppableProvided, _droppableSnapshot) => (
              <div ref={droppableProvided.innerRef} style={{ display: 'flex', flexDirection: 'column' }}>
                {list.map((tag, index) => (
                  <Draggable key={String(index)} draggableId={String(index)} index={index}>
                    {(draggableProvided, _draggableSnapshot) => (
                      <div ref={draggableProvided.innerRef} {...draggableProvided.draggableProps}>
                        <div className="draggableItem">
                          <div className="dragHandle" {...draggableProvided.dragHandleProps}>
                            <ThreeBarsIcon />
                          </div>
                          {renderItem(tag, index)}
                        </div>
                      </div>
                    )}
                  </Draggable>
                ))}
                {droppableProvided.placeholder}
              </div>
            )}
          </Droppable>
        </DragDropContext>
      </div>
    );
  }
}
