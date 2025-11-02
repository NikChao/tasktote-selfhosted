import { List, ListItem, Sheet, Typography } from "@mui/joy";
import {
  GroceryItem,
  LayoutBlock,
  ScheduledDays,
} from "../../services/grocery-service";
import { TaskRow } from "./components/task";
import { GroceryItemRow } from "./components/grocery-item";

enum Values {
  Unknown = "unknown",
}

const containerStyles = {
  width: "100%",
  bgcolor: "background.paper",
};

export function GroceryList({
  groceries,
  layout,
  schedule,
  checkGroceryItem,
  saveTaskScheduledDays,
}: {
  groceries: GroceryItem[];
  layout: LayoutBlock[];
  schedule: { taskId: string; date: string }[];
  checkGroceryItem(id: string): void;
  saveTaskScheduledDays: (id: string, scheduledDays: ScheduledDays) => void;
}) {
  if (!groceries?.length) {
    return (
      <Sheet sx={{ backgroundColor: "white", paddingTop: "16px" }}>
        <Typography
          level="body-sm"
          textColor="neutral.400"
          fontStyle="italic"
          textAlign="center"
        >
          No items yet...
        </Typography>
      </Sheet>
    );
  }

  return (
    <List sx={containerStyles}>
      {layout.map(({ type, value }) => {
        if (value === Values.Unknown) return null;

        if (type === "Text") {
          return (
            <ListItem
              sx={{ width: "100%", height: "48px" }}
              key={value}
            >
              <Typography level="h4">{value}</Typography>
            </ListItem>
          );
        }

        const item = groceries.find(({ id }) => id === value);
        if (!item) {
          return null;
        }

        if (item.kind === "Grocery") {
          return (
            <GroceryItemRow
              key={item.id}
              groceryItem={item}
              checkGroceryItem={checkGroceryItem}
            />
          );
        }

        if (item.kind === "Task") {
          return (
            <TaskRow
              key={item.id}
              task={item}
              dates={schedule.filter(({ taskId }) => taskId === item.id).map((
                { date },
              ) => new Date(date))}
              checkTask={checkGroceryItem}
              saveTaskScheduledDays={saveTaskScheduledDays}
            />
          );
        }

        return item.kind satisfies never;
      })}
    </List>
  );
}
