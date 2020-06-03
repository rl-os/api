import { createContext, useContext } from 'react';
import { Api } from '../api';
import { Config } from './config';
import { Auth } from './auth';
import { observable } from 'mobx';

export class Store {
  public readonly config: Config;
  public readonly api: Api;

  public readonly auth: Auth;

  constructor() {
    this.config = observable(new Config(this));
    this.auth = observable(new Auth(this));

    this.api = observable(new Api(this));
  }
}

export const StoreContext = createContext<Store | null>(null);
export const useStore = (): Store => useContext(StoreContext)!;
