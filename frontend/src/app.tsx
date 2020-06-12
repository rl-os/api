import React, { useRef } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { observer } from 'mobx-react-lite';
import { hot } from 'react-hot-loader/root';
import { Store, StoreContext } from './store';
import { useAsync } from './utils/hooks';
import { SecureRoute } from './components/secure-route';

import { HomePage } from './pages/home';
import { LoginPage } from "./pages/auth/login";

const content = (store: Store) => {
  return <StoreContext.Provider value={store}>
    <BrowserRouter>
      <SecureRoute exact={true} path="/" component={HomePage}/>
      <SecureRoute exact={true} path="/auth/login" component={LoginPage}/>
    </BrowserRouter>
  </StoreContext.Provider>;
};

const App = observer(() => {
  const storeRef = useRef<Store>();
  let store: Store;

  if (storeRef.current === undefined || Object.getPrototypeOf(storeRef.current) !== Store.prototype)
    storeRef.current = store = new Store();
  else store = storeRef.current;

  return useAsync({
    dependencies: [store],
    init: async () => {
      await store.auth.init();
    },
    loading: () => <div>Loading</div>,
    failed: (e) => <div>{e.toString()}</div>,
    success: () => content(store),
  });
});

export default process.env.NODE_ENV === 'development' ? hot(App) : App;
