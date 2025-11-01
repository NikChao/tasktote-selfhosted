import { observer } from "mobx-react-lite";
import { GroceryScreen } from "./screen";
import { UserStore } from "../../store/user-store";
import { GroceryListStore } from "./store";

export function createGroceryScreen(
  store: GroceryListStore,
  userStore: UserStore,
  endIcons: React.ReactNode
) {
  return observer(() => (
    <GroceryScreen
      mode={store.mode}
      selectedStores={store.selectedStores}
      toggleStore={store.toggleStore}
      endIcons={endIcons}
      groceryItemText={store.groceryItemText}
      isUpdating={store.isUpdating}
      groceryList={store.filteredGroceryList}
      householdId={userStore.effectiveHouseholdId}
      setMode={store.setMode}
      handleGroceryItemTextChange={store.handleGroceryItemTextChange}
      checkGroceryItem={store.checkGroceryItem}
      clearCheckedItems={store.clearCheckedItems}
      createItem={store.createItem}
      fetchGroceries={store.fetchGroceriesIfTabFocussed}
      initializeGroceryList={store.initializeGroceryList}
      saveScheduledDays={store.saveTaskScheduledDays}
    />
  ));
}
