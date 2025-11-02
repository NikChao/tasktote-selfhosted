import {
  Box,
  Checkbox,
  DialogTitle,
  IconButton,
  ListItem,
  ListItemButton,
  Modal,
  ModalClose,
  ModalDialog,
  Typography,
} from "@mui/joy";
import {
  DEFAULT_SCHEDULED_DAYS,
  GroceryItem as Task,
  ScheduledDays,
} from "../../../services/grocery-service";
import { MouseEvent, useState } from "react";
import { CalendarMonth } from "@mui/icons-material";
import { Calendar } from "./calendar";

interface TaskRowProps {
  task: Task;
  dates: Date[];
  checkTask: (id: string) => void;
  saveTaskScheduledDays: (id: string, scheduledDays: ScheduledDays) => void;
}

export function TaskRow(
  { task, checkTask, saveTaskScheduledDays, dates }: TaskRowProps,
) {
  const [calendarOpen, setCalendarOpen] = useState(false);

  function handleCalendarClick(e: MouseEvent<HTMLButtonElement>) {
    e.stopPropagation();
    setCalendarOpen(true);
  }

  return (
    <>
      <ListItem
        sx={{ width: "100%", height: "48px" }}
        onClick={() => {
          checkTask(task.id);
        }}
      >
        <ListItemButton
          sx={{
            display: "flex",
            justifyContent: "space-between",
            borderRadius: "12px",
          }}
        >
          <Typography>{task.name}</Typography>

          <Box display="flex" alignItems="center" gap="8px">
            <IconButton onClick={handleCalendarClick}>
              <CalendarMonth color="primary" />
            </IconButton>
            <Checkbox checked={task.checked} />
          </Box>
        </ListItemButton>
      </ListItem>
      <Modal
        open={calendarOpen}
        onClose={() => {
          setCalendarOpen(false);
        }}
      >
        <ModalDialog layout="fullscreen">
          <ModalClose />
          <DialogTitle>Schedule task</DialogTitle>
          <Box width="100%" height="100%" display="flex" alignItems="center">
            <Box flex="1">
              <Calendar
                scheduledDates={dates}
                onSave={(newDays) => {
                  saveTaskScheduledDays(task.id, newDays);
                  setCalendarOpen(false);
                }}
              />
            </Box>
          </Box>
        </ModalDialog>
      </Modal>
    </>
  );
}
