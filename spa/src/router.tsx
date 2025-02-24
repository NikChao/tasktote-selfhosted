import {
  createRootRoute,
  createRoute,
  createRouter,
  Outlet,
  useNavigate,
} from "@tanstack/react-router";
import { createGroceryScreen } from "./features/grocery/create";
import { UserStore } from "./store/user-store";
import { CreateAndInviteToHousehold as CreateAndInviteToHouseholdImpl } from "./features/household/invite-to-household";
import { UserService } from "./services/user-service";
import { ApiService } from "./services/api-service";
import { GroceryService } from "./services/grocery-service";
import { HouseholdService } from "./services/household-service";
import { observer } from "mobx-react-lite";
import { Box, CircularProgress } from "@mui/joy";
import { AutoCompleteGroceries } from "./features/end-buttons/autocomplete-groceries";
import { GroceryListStore } from "./features/grocery/store";

const apiService = new ApiService();
const groceryService = new GroceryService(apiService);
const userService = new UserService(apiService);
const householdService = new HouseholdService(apiService);
const userStore = new UserStore(userService, householdService);
const groceryStore = new GroceryListStore(groceryService, userStore);

const CreateAndInviteToHousehold = observer(() => {
  if (!userStore.userId) {
    return <CircularProgress size="md" />;
  }

  return (
    <CreateAndInviteToHouseholdImpl
      householdId={userStore.effectiveHouseholdId}
      isLoading={userStore.isLoading}
      leaveHousehold={userStore.leaveHousehold}
      joinHousehold={userStore.joinHousehold}
      createAndJoinHousehold={userStore.createAndJoinHousehold}
    />
  );
});

const EndIcons = observer(() => (
  <Box display="flex" alignItems="center">
    <AutoCompleteGroceries
      isLoading={groceryStore.isFetching}
      householdId={userStore.effectiveHouseholdId}
      isMagicEnabled={groceryStore.magicEnabled}
      magic={groceryStore.magic}
    />
    <CreateAndInviteToHousehold />
  </Box>
));

const endIcons = <EndIcons />;

const GroceryScreen = createGroceryScreen(groceryStore, userStore, endIcons);

const rootRoute = createRootRoute({
  component: () => <Outlet />,
});

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: GroceryScreen,
});

const joinHouseholdRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/households/join/$householdId",
  component: function JoinHousehold() {
    /** @ts-ignore */
    const { householdId } = joinHouseholdRoute.useParams();
    const navigate = useNavigate();

    userStore.joinHousehold(householdId).then(() => navigate({ to: "/" }));

    return null;
  },
});

/** @ts-ignore */
const routeTree = rootRoute.addChildren([
  indexRoute,
  joinHouseholdRoute,
]);

export const router = createRouter({ routeTree });
