import { createContext, useContext } from 'react';
import { Api } from '../api';
import { Config } from './config';
import { Auth } from './auth';
import { observable } from 'mobx';

export class Store {
  @observable
  public readonly config: Config;
  @observable
  public readonly api: Api;
  @observable
  public readonly auth: Auth;

  constructor() {
    this.config = new Config(this);
    this.auth = new Auth(this);

    this.api = new Api(this);
  }
}

export const StoreContext = createContext<Store | null>(null);
export const useStore = (): Store => useContext(StoreContext)!;
