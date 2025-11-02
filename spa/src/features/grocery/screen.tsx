import { ChangeEvent, useEffect } from "react";
import {
  Box,
  Button,
  Input,
  Stack,
  Typography,
  TypographyProps,
} from "@mui/joy";
import { FormEvent } from "react";
import {
  GroceryItemKind,
  GroceryList as GroceryListModel,
  ScheduledDays,
  StoreName,
} from "../../services/grocery-service";
import { poll } from "../../utils/poll";
import { GroceryList } from "./grocery-list";

const POLL_INTERVAL = 8000;

interface GroceryScreenProps {
  mode: GroceryItemKind;
  groceryItemText: string;
  isUpdating: boolean;
  householdId: string;
  endIcons: React.ReactNode;
  groceryList: GroceryListModel;
  selectedStores: StoreName[];
  schedule: { taskId: string; date: string }[];
  setMode: (mode: GroceryItemKind) => void;
  handleGroceryItemTextChange(e: ChangeEvent<HTMLInputElement>): void;
  checkGroceryItem(id: string): void;
  clearCheckedItems(householdId: string): void;
  createItem(
    householdId: string,
  ): (e: FormEvent<HTMLFormElement>) => void;
  initializeGroceryList(householdId: string): void;
  fetchGroceries(householdId: string): void;
  toggleStore(storeName: StoreName): void;
  saveScheduledDays(id: string, scheduledDays: ScheduledDays): void;
}

export function GroceryScreen({
  mode,
  groceryItemText,
  groceryList,
  isUpdating,
  endIcons,
  householdId,
  schedule,
  setMode,
  handleGroceryItemTextChange,
  checkGroceryItem,
  clearCheckedItems,
  createItem,
  initializeGroceryList,
  fetchGroceries,
  saveScheduledDays,
}: GroceryScreenProps) {
  useEffect(() => {
    initializeGroceryList(householdId);
    const interval = poll(() => fetchGroceries(householdId), POLL_INTERVAL);

    return () => {
      clearInterval(interval);
    };
  }, [householdId, initializeGroceryList, fetchGroceries]);

  function getTitleProps(kind: GroceryItemKind): Partial<TypographyProps> {
    const baseProps: Partial<TypographyProps> = {
      sx: {
        cursor: "pointer",
      },
      style: {},
      onClick: () => setMode(kind),
    };

    if (kind !== mode) {
      baseProps.style!.opacity = 0.6;
      return baseProps;
    }

    baseProps.fontStyle = "italic";
    return baseProps;
  }

  return (
    <Stack
      height="100%"
      display="flex"
      justifyContent="space-between"
      padding="24px"
      boxSizing="border-box"
    >
      <Stack maxHeight="90%">
        <Stack>
          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
          >
            <Box display="flex" gap="8px">
              <Typography
                level="h3"
                color="primary"
                {...getTitleProps("Grocery")}
              >
                Groceries
              </Typography>
              <Typography level="h3" color="success" {...getTitleProps("Task")}>
                Tasks
              </Typography>
            </Box>
            {endIcons}
          </Box>
          <Box
            component="form"
            onSubmit={createItem(householdId)}
            pt="8px"
            boxSizing="border-box"
          >
            <Input
              value={groceryItemText}
              onChange={handleGroceryItemTextChange}
              placeholder={mode === "Grocery"
                ? "Add a grocery item or recipe url!"
                : "Add some tasks!"}
            />
          </Box>
        </Stack>
        <Box overflow="scroll" width="100%">
          <GroceryList
            groceries={groceryList.items}
            layout={groceryList.layout}
            checkGroceryItem={checkGroceryItem}
            saveTaskScheduledDays={saveScheduledDays}
            schedule={schedule}
          />
        </Box>
      </Stack>
      <Stack>
        <Button
          color="danger"
          variant="soft"
          fullWidth
          onClick={() => clearCheckedItems(householdId)}
          disabled={isUpdating}
          loading={isUpdating}
        >
          delete checked
        </Button>
      </Stack>
    </Stack>
  );
}
