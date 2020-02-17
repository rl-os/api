import { createContext, useContext } from 'react';

export class Store {
  constructor() {}
}

export const StoreContext = createContext<Store | null>(null);
export const useStore = (): Store => useContext(StoreContext)!;
