import { ChevronLeft, ChevronRight } from "@mui/icons-material";
import { Box, Button, IconButton, Typography } from "@mui/joy";
import { useState } from "react";
import {
  DAYS,
  DEFAULT_SCHEDULED_DAYS,
  MONTHS,
  ScheduledDays,
} from "../../../services/grocery-service";

interface CalendarProps {
  scheduledDates: Date[];
  onSave: (newDays: ScheduledDays) => void;
}

function borderProps(week: number, day: number) {
  const borderLeft = day === 0 ? "1px solid black" : undefined;
  const borderRight = "1px solid black";
  const borderTop = week === 0 ? "1px solid black" : undefined;
  const borderBottom = "1px solid black";

  const borderTopLeftRadius = (week === 0 && day === 0) ? "12px" : undefined;
  const borderTopRightRadius = week === 0 && day === 6 ? "12px" : undefined;
  const borderBottomLeftRadius = (week === 4 && day === 0) ? "12px" : undefined;
  const borderBottomRightRadius = week === 4 && day === 6 ? "12px" : undefined;

  return {
    borderLeft,
    borderRight,
    borderTop,
    borderBottom,

    borderTopLeftRadius,
    borderTopRightRadius,
    borderBottomLeftRadius,
    borderBottomRightRadius,
  };
}

export function Calendar({
  scheduledDates,
  onSave,
}: CalendarProps) {
  const parsedDates = scheduledDates.map((date) => {
    return {
      date: date.getDate(),
      month: Object.keys(DEFAULT_SCHEDULED_DAYS)[date.getMonth()]!,
    };
  }).reduce((prev, accum) => {
    prev[accum.month as keyof ScheduledDays].push(accum.date);
    return prev;
  }, DEFAULT_SCHEDULED_DAYS);

  const [scheduledDays, setScheduledDays] = useState<ScheduledDays>(
    parsedDates,
  );
  const [monthIndex, setMonthIndex] = useState(new Date().getMonth());

  function lastMonth() {
    setMonthIndex(monthIndex === 0 ? 11 : monthIndex - 1);
  }

  function nextMonth() {
    setMonthIndex(monthIndex === 11 ? 0 : monthIndex + 1);
  }

  const monthName = Object.keys(MONTHS)[monthIndex] as keyof typeof MONTHS;
  const month = MONTHS[monthName];
  const weekCount = Math.floor(month.days / 7) + 1;
  const weeks = new Array(weekCount).fill(0).map((_, i) => i);

  function handleClickDate(day: number) {
    if (scheduledDays[monthName].includes(day)) {
      setScheduledDays({
        ...scheduledDays,
        [monthName]: scheduledDays[monthName].filter((d) => d !== day),
      });
    } else {
      setScheduledDays({
        ...scheduledDays,
        [monthName]: [...scheduledDays[monthName], day],
      });
    }
  }

  return (
    <Box onClick={(e) => e.preventDefault()}>
      <Box>
        <Box display="flex" justifyContent="space-around">
          {DAYS.map((day, index) => {
            return (
              <Box key={index}>
                {day}
              </Box>
            );
          })}
        </Box>
        {weeks.map((week) => {
          return (
            <Box display="flex" key={week}>
              {DAYS.map((_, dayOfWeek) => {
                const dayNum = 7 * week + dayOfWeek;
                if (dayNum > month.days) {
                  return (
                    <Box
                      key={dayNum}
                      flex="1"
                      sx={{ ...borderProps(week, dayOfWeek) }}
                    />
                  );
                }

                return (
                  <Box
                    key={dayNum}
                    component="button"
                    flex="1"
                    minHeight="80px"
                    display="flex"
                    flexDirection="column"
                    alignItems="center"
                    padding="0"
                    margin="0"
                    bgcolor="transparent"
                    border="none"
                    sx={{ ...borderProps(week, dayOfWeek) }}
                    onClick={() => handleClickDate(dayNum)}
                  >
                    <Typography>{dayNum + 1}</Typography>
                    {scheduledDays[monthName].includes(dayNum)
                      ? (
                        <Box
                          borderRadius="100%"
                          width="8px"
                          height="8px"
                          bgcolor="lightseagreen"
                        />
                      )
                      : null}
                  </Box>
                );
              })}
            </Box>
          );
        })}
      </Box>
      <Box
        display="flex"
        flexDirection="column"
        alignItems="center"
        justifyContent="center"
      >
        <Box display="flex" justifyContent="center" alignItems="center">
          <IconButton onClick={lastMonth}>
            <ChevronLeft />
          </IconButton>
          <Typography>{monthName}</Typography>
          <IconButton onClick={nextMonth}>
            <ChevronRight />
          </IconButton>
        </Box>

        <Box paddingTop="32px" width="100%">
          <Button
            fullWidth
            color="primary"
            variant="soft"
            onClick={(e) => {
              e.stopPropagation();
              e.preventDefault();
              onSave(scheduledDays);
            }}
          >
            Save
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
