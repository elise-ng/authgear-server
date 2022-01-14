/* global JSX */
import React, { useMemo, useCallback, useContext, useState } from "react";
import { FormattedMessage, Context } from "@oursky/react-messageformat";
import {
  DetailsList,
  DetailsHeader,
  DetailsRow,
  DirectionalHint,
  Dropdown,
  Dialog,
  DialogFooter,
  PrimaryButton,
  DefaultButton,
  IconButton,
  SelectionMode,
  IColumn,
  IDropdownOption,
  IDialogContentProps,
  IDetailsHeaderProps,
  IDetailsRowProps,
  IDetailsColumnRenderTooltipProps,
  IRenderFunction,
  IIconProps,
  IDragDropEvents,
  Icon,
} from "@fluentui/react";
import LabelWithTooltip from "./LabelWithTooltip";
import {
  UserProfileAttributesAccessControl,
  AccessControlLevelString,
} from "./types";
import { parseJSONPointer } from "./util/jsonpointer";
import styles from "./UserProfileAttributesList.module.scss";

export type UserProfileAttributesListAccessControlAdjustment = [
  keyof UserProfileAttributesAccessControl,
  AccessControlLevelString
];

export interface UserProfileAttributesListItem {
  pointer: string;
  access_control: UserProfileAttributesAccessControl;
}

export interface ItemComponentProps<T> {
  className: string;
  item: T;
}

export interface UserProfileAttributesListProps<
  T extends UserProfileAttributesListItem
> {
  items: T[];
  ItemComponent: React.ComponentType<ItemComponentProps<T>>;
  onChangeItems: (items: T[]) => void;
  onEditButtonClick?: (index: number) => void;
  onReorderItems?: (items: T[]) => void;
}

export interface UserProfileAttributesListPendingUpdate {
  index: number;
  key: keyof UserProfileAttributesAccessControl;
  mainAdjustment: UserProfileAttributesListAccessControlAdjustment;
  otherAdjustments: UserProfileAttributesListAccessControlAdjustment[];
}

const EDIT_BUTTON_ICON_PROPS: IIconProps = {
  iconName: "Edit",
};

function intOfAccessControlLevelString(
  level: AccessControlLevelString
): number {
  switch (level) {
    case "hidden":
      return 1;
    case "readonly":
      return 2;
    case "readwrite":
      return 3;
    default:
      throw new Error("unknown value: " + String(level));
  }
}

function accessControlLevelStringOfInt(
  value: number
): AccessControlLevelString {
  switch (value) {
    case 1:
      return "hidden";
    case 2:
      return "readonly";
    case 3:
      return "readwrite";
  }
  throw new Error("unknown value: " + String(value));
}

function adjustAccessControl(
  accessControl: UserProfileAttributesAccessControl,
  target: keyof UserProfileAttributesAccessControl,
  refValue: AccessControlLevelString
): UserProfileAttributesListAccessControlAdjustment | undefined {
  const targetLevelInt = intOfAccessControlLevelString(accessControl[target]);
  const refLevelInt = intOfAccessControlLevelString(refValue);
  if (targetLevelInt <= refLevelInt) {
    return undefined;
  }

  return [target, accessControlLevelStringOfInt(refLevelInt)];
}

function makeUpdate<T extends UserProfileAttributesListItem>(
  prevItems: T[],
  index: number,
  key: keyof UserProfileAttributesAccessControl,
  newValue: AccessControlLevelString
): UserProfileAttributesListPendingUpdate {
  const accessControl = prevItems[index].access_control;
  const mainAdjustment: UserProfileAttributesListAccessControlAdjustment = [
    key,
    newValue,
  ];

  const adjustments: ReturnType<typeof adjustAccessControl>[] = [];
  switch (key) {
    case "end_user":
      break;
    case "bearer": {
      if (newValue === "hidden") {
        adjustments.push(
          adjustAccessControl(accessControl, "end_user", newValue)
        );
      }
      break;
    }
    case "portal_ui": {
      adjustments.push(adjustAccessControl(accessControl, "bearer", newValue));
      adjustments.push(
        adjustAccessControl(accessControl, "end_user", newValue)
      );
      break;
    }
  }

  const otherAdjustments: UserProfileAttributesListAccessControlAdjustment[] =
    adjustments.filter(
      (a): a is UserProfileAttributesListAccessControlAdjustment => a != null
    );

  return {
    index,
    key,
    mainAdjustment,
    otherAdjustments,
  };
}

function applyUpdate<T extends UserProfileAttributesListItem>(
  prevItems: T[],
  update: UserProfileAttributesListPendingUpdate
): T[] {
  const { index, mainAdjustment, otherAdjustments } = update;
  let accessControl = prevItems[index].access_control;
  const adjustments = [mainAdjustment, ...otherAdjustments];

  for (const adjustment of adjustments) {
    accessControl = {
      ...accessControl,
      [adjustment[0]]: adjustment[1],
    };
  }

  const newItems = [...prevItems];
  newItems[index] = {
    ...newItems[index],
    access_control: accessControl,
  };

  return newItems;
}

function UserProfileAttributesList<T extends UserProfileAttributesListItem>(
  props: UserProfileAttributesListProps<T>
): React.ReactElement<any, any> | null {
  const {
    items,
    onChangeItems,
    ItemComponent,
    onEditButtonClick,
    onReorderItems,
  } = props;
  const { renderToString } = useContext(Context);
  const [pendingUpdate, setPendingUpdate] = useState<
    UserProfileAttributesListPendingUpdate | undefined
  >();
  const [dndIndex, setDNDIndex] = useState<number | undefined>(undefined);

  const reorder = useCallback(
    (index: number, item: T) => {
      const itemsWithoutIndex = [
        ...items.slice(0, index),
        ...items.slice(index + 1),
      ];
      const insertIndex = items.indexOf(item);
      if (insertIndex >= 0) {
        itemsWithoutIndex.splice(insertIndex, 0, items[index]);
        onReorderItems?.(itemsWithoutIndex);
      }
    },
    [items, onReorderItems]
  );

  const dragDropEvents: IDragDropEvents = useMemo(() => {
    return {
      canDrop: () => true,
      canDrag: () => true,
      onDragEnter: () => styles.onDragEnter,
      onDragLeave: () => {},
      onDragStart: (_item?: T, index?: number) => {
        if (index != null) {
          setDNDIndex(index);
        }
      },
      onDragEnd: (_item?: T) => {
        setDNDIndex(undefined);
      },
      onDrop: (item?: T) => {
        if (dndIndex != null && item != null) {
          reorder(dndIndex, item);
        }
      },
    };
  }, [reorder, dndIndex]);

  const onClickConfirmPendingUpdate = useCallback(
    (e: React.MouseEvent<unknown>) => {
      e.preventDefault();
      e.stopPropagation();

      if (pendingUpdate != null) {
        const newItems = applyUpdate(items, pendingUpdate);
        setPendingUpdate(undefined);
        onChangeItems(newItems);
      }
    },
    [items, onChangeItems, pendingUpdate]
  );

  const onDismissPendingUpdateDialog = useCallback(() => {
    setPendingUpdate(undefined);
  }, []);

  // title and subText are typed as string but they can actually be any JSX.Element.
  // @ts-expect-error
  const pendingUpdateDialogContentProps: IDialogContentProps = useMemo(() => {
    if (pendingUpdate == null) {
      return {
        title: "",
        subText: "",
      };
    }

    const { index } = pendingUpdate;

    const pointer = items[index].pointer;
    const fieldName = parseJSONPointer(pointer)[0];

    const formattedLevel = renderToString(
      "user-profile.access-control-level." + pendingUpdate.mainAdjustment[1]
    );

    const affected =
      pendingUpdate.otherAdjustments.length === 1
        ? pendingUpdate.otherAdjustments[0][0]
        : "other";

    return {
      title: (
        <FormattedMessage
          id="UserProfileAttributesList.dialog.title.pending-update"
          values={{
            fieldName,
            party: pendingUpdate.mainAdjustment[0],
          }}
        />
      ),
      subText: (
        <FormattedMessage
          id="UserProfileAttributesList.dialog.description.pending-update"
          values={{
            fieldName,
            affected,
            party: pendingUpdate.mainAdjustment[0],
            level: formattedLevel,
          }}
        />
      ),
    };
  }, [renderToString, pendingUpdate, items]);

  const makeDropdownOnChange = useCallback(
    (index: number, key: keyof UserProfileAttributesAccessControl) => {
      return (
        _e: React.FormEvent<unknown>,
        option?: IDropdownOption<AccessControlLevelString>,
        _index?: number
      ) => {
        if (option == null) {
          return;
        }

        const pendingUpdate = makeUpdate(
          items,
          index,
          key,
          option.key as AccessControlLevelString
        );

        if (pendingUpdate.otherAdjustments.length !== 0) {
          setPendingUpdate(pendingUpdate);
          return;
        }

        const newItems = applyUpdate(items, pendingUpdate);
        onChangeItems(newItems);
      };
    },
    [items, onChangeItems]
  );

  const makeRenderDropdown = useCallback(
    (key: keyof UserProfileAttributesAccessControl) => {
      return (
        item?: UserProfileAttributesListItem,
        index?: number,
        _column?: IColumn
      ) => {
        if (item == null || index == null) {
          return null;
        }

        const optionHidden: IDropdownOption = {
          key: "hidden",
          text: renderToString("user-profile.access-control-level.hidden"),
        };

        const optionReadonly: IDropdownOption = {
          key: "readonly",
          text: renderToString("user-profile.access-control-level.readonly"),
        };

        const optionReadwrite: IDropdownOption = {
          key: "readwrite",
          text: renderToString("user-profile.access-control-level.readwrite"),
        };

        let options: IDropdownOption<AccessControlLevelString>[] = [];
        let selectedKey: string | undefined;
        switch (key) {
          case "portal_ui":
            options = [optionHidden, optionReadonly, optionReadwrite];
            selectedKey = item.access_control.portal_ui;
            break;
          case "bearer":
            options = [optionHidden, optionReadonly];
            if (item.access_control.portal_ui === "hidden") {
              optionReadonly.disabled = true;
            }
            selectedKey = item.access_control.bearer;
            break;
          case "end_user":
            options = [optionHidden, optionReadonly, optionReadwrite];
            if (item.access_control.bearer === "hidden") {
              optionReadwrite.disabled = true;
              optionReadonly.disabled = true;
            }
            if (item.access_control.portal_ui === "hidden") {
              optionReadwrite.disabled = true;
              optionReadonly.disabled = true;
            }
            if (item.access_control.portal_ui === "readonly") {
              optionReadwrite.disabled = true;
            }
            selectedKey = item.access_control.end_user;
            break;
        }

        const disabledOptionCount = options.reduce((a, b) => {
          return a + (b.disabled === true ? 1 : 0);
        }, 0);
        const dropdownIsDisabled = options.length - disabledOptionCount <= 1;

        return (
          <Dropdown
            options={options}
            selectedKey={selectedKey}
            disabled={dropdownIsDisabled}
            onChange={makeDropdownOnChange(index, key)}
          />
        );
      };
    },
    [renderToString, makeDropdownOnChange]
  );

  const onRenderPointer = useCallback(
    (item?: T, _index?: number, _column?: IColumn) => {
      if (item == null) {
        return null;
      }
      return <ItemComponent className="" item={item} />;
    },
    [ItemComponent]
  );

  const onRenderEditButton = useCallback(
    (
      _item?: UserProfileAttributesListItem,
      index?: number,
      _column?: IColumn
    ) => {
      if (index == null) {
        return null;
      }
      const onClick = (e: React.MouseEvent<unknown>) => {
        e.preventDefault();
        e.stopPropagation();
        onEditButtonClick?.(index);
      };
      return (
        <IconButton
          iconProps={EDIT_BUTTON_ICON_PROPS}
          title={renderToString("edit")}
          ariaLabel={renderToString("edit")}
          onClick={onClick}
        />
      );
    },
    [onEditButtonClick, renderToString]
  );

  const onRenderReorderHandle = useCallback(() => {
    return (
      <div className={styles.reorderHandle}>
        <Icon iconName="GlobalNavButton" />
      </div>
    );
  }, []);

  const columns: IColumn[] = useMemo(() => {
    const columns: IColumn[] = [
      {
        key: "pointer",
        minWidth: 200,
        name: renderToString(
          "UserProfileAttributesList.header.label.attribute-name"
        ),
        onRender: onRenderPointer,
        isMultiline: true,
      },
      {
        key: "portal_ui",
        minWidth: 200,
        maxWidth: 200,
        name: "",
        onRender: makeRenderDropdown("portal_ui"),
      },
      {
        key: "bearer",
        minWidth: 200,
        maxWidth: 200,
        name: "",
        onRender: makeRenderDropdown("bearer"),
      },
      {
        key: "end_user",
        minWidth: 200,
        maxWidth: 200,
        name: "",
        onRender: makeRenderDropdown("end_user"),
      },
    ];
    if (onEditButtonClick != null) {
      columns.push({
        key: "edit",
        minWidth: 24,
        maxWidth: 24,
        name: "",
        onRender: onRenderEditButton,
      });
    }
    if (onReorderItems != null) {
      columns.push({
        key: "reorder",
        minWidth: 24,
        maxWidth: 24,
        name: "",
        onRender: onRenderReorderHandle,
      });
    }
    return columns;
  }, [
    onEditButtonClick,
    onReorderItems,
    renderToString,
    makeRenderDropdown,
    onRenderPointer,
    onRenderEditButton,
    onRenderReorderHandle,
  ]);

  const onRenderColumnHeaderTooltip: IRenderFunction<IDetailsColumnRenderTooltipProps> =
    useCallback(
      (
        props?: IDetailsColumnRenderTooltipProps,
        defaultRender?: (
          props: IDetailsColumnRenderTooltipProps
        ) => JSX.Element | null
      ) => {
        if (props == null || defaultRender == null) {
          return null;
        }
        if (props.column == null) {
          return null;
        }
        if (
          props.column.key === "portal_ui" ||
          props.column.key === "bearer" ||
          props.column.key === "end_user"
        ) {
          return (
            <LabelWithTooltip
              labelId={
                "UserProfileAttributesList.header.label." + props.column.key
              }
              tooltipMessageId={
                "UserProfileAttributesList.header.tooltip." + props.column.key
              }
              directionalHint={DirectionalHint.topCenter}
            />
          );
        }
        return defaultRender(props);
      },
      []
    );

  const onRenderDetailsHeader: IRenderFunction<IDetailsHeaderProps> =
    useCallback(
      (props?: IDetailsHeaderProps) => {
        if (props == null) {
          return null;
        }
        return (
          <DetailsHeader
            {...props}
            onRenderColumnHeaderTooltip={onRenderColumnHeaderTooltip}
          />
        );
      },
      [onRenderColumnHeaderTooltip]
    );

  const onRenderRow: IRenderFunction<IDetailsRowProps> = useCallback(
    (props?: IDetailsRowProps) => {
      if (props == null) {
        return null;
      }
      let className = "";
      const { itemIndex } = props;
      if (dndIndex != null) {
        if (itemIndex < dndIndex) {
          className = styles.before;
        } else if (itemIndex > dndIndex) {
          className = styles.after;
        }
      }
      return <DetailsRow {...props} className={className} />;
    },
    [dndIndex]
  );

  return (
    <>
      <DetailsList
        columns={columns}
        items={items}
        selectionMode={SelectionMode.none}
        onRenderDetailsHeader={onRenderDetailsHeader}
        onRenderRow={onRenderRow}
        dragDropEvents={onReorderItems != null ? dragDropEvents : undefined}
      />
      <Dialog
        hidden={pendingUpdate == null}
        onDismiss={onDismissPendingUpdateDialog}
        dialogContentProps={pendingUpdateDialogContentProps}
      >
        <DialogFooter>
          <PrimaryButton onClick={onClickConfirmPendingUpdate}>
            <FormattedMessage id="confirm" />
          </PrimaryButton>
          <DefaultButton onClick={onDismissPendingUpdateDialog}>
            <FormattedMessage id="cancel" />
          </DefaultButton>
        </DialogFooter>
      </Dialog>
    </>
  );
}

export default UserProfileAttributesList;
