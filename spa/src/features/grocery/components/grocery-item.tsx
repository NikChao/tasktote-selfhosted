import { Checkbox, ListItem, ListItemButton, Typography } from "@mui/joy";
import { GroceryItem } from "../../../services/grocery-service";

interface GroceryItemRowProps {
  groceryItem: GroceryItem;
  checkGroceryItem: (id: string) => void;
}

export function GroceryItemRow({ groceryItem, checkGroceryItem }: GroceryItemRowProps) {
  return (
    <ListItem
      sx={{ width: "100%", height: "48px" }}
      onClick={() => {
        checkGroceryItem(groceryItem.id);
      }}
    >
      <ListItemButton
        sx={{ display: "flex", justifyContent: "space-between", borderRadius: "12px" }}
      >
        <Typography>{groceryItem.name}</Typography>
        <Checkbox checked={groceryItem.checked} />
      </ListItemButton>
    </ListItem>
  )
}
