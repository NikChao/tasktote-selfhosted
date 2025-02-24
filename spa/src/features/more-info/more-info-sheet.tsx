import { Button, Drawer, Stack, Typography } from "@mui/joy";
import { useState } from "react";
import { ALL_STORES, StoreName } from "../../services/grocery-service";
import { Checkbox } from "@mui/joy";
import { observer } from "mobx-react-lite";

interface MoreInfoSheetProps {
  selectedStores: StoreName[];
  toggleStore(storeName: StoreName): void;
}

export const MoreInfoSheet = observer(
  ({ selectedStores, toggleStore }: MoreInfoSheetProps) => {
    const [isOpen, setIsOpen] = useState(false);

    function open() {
      setIsOpen(true);
    }

    function close() {
      setIsOpen(false);
    }

    return (
      <>
        <Button color="primary" variant="plain" onClick={open}>
          more info
        </Button>
        <Drawer open={isOpen} onClose={close} anchor="bottom">
          <Stack p="16px" height="70vh" gap="16px">
            <Stack>
              <Typography level="body-md" fontWeight="600">
                More info
              </Typography>
            </Stack>
            <Stack gap={2}>
              <Typography level="body-sm" fontWeight="600">
                What stores do you shop at?
              </Typography>
              {ALL_STORES.map((store) => (
                <Checkbox
                  key={store}
                  name={store}
                  label={store}
                  onChange={() => toggleStore(store)}
                  checked={selectedStores.includes(store)}
                />
              ))}
            </Stack>
          </Stack>
        </Drawer>
      </>
    );
  },
);
