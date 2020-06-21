import { observable } from "mobx";
import { Store } from ".";
import { DetailedUser } from "../models/detailed_user";

const LOCALSTORAGE_KEY = 'user_auth';

export class Auth {
  @observable
  public currentUser: DetailedUser | null = null;

  @observable
  public accessToken: string | null = null;

  @observable
  public refreshToken: string | null = null;

  constructor(private readonly store: Store) {}

  /**
   * @async
   * @description Calling before load all application
   */
  public async init() {
    if (this.currentUser !== null) {
      return;
    }

    const data = JSON.parse(
      localStorage.getItem(LOCALSTORAGE_KEY)!,
    );
    if (data) {
      this.accessToken = data.accessToken || null;
      this.refreshToken = data.refreshToken || null;
    }

    try {
      this.currentUser = await this.store.api.user.me();
    } catch (e) {
      if (e.isAxiosError && e.response?.status === 401) {
        this.currentUser = null;
        return;
      }

      throw e;
    }
  }

  protected save() {
    localStorage.setItem(LOCALSTORAGE_KEY, JSON.stringify({
      accessToken: this.accessToken,
      refreshToken: this.refreshToken
    }))
  }

  public reset() {
    this.currentUser = null;
  }

  public async login(username: string, pwd: string): Promise<void> {
    const token = await this.store.api.auth.loginByPwd(username, pwd);

    this.accessToken = token.access_token;
    this.refreshToken = token.refresh_token;

    this.currentUser = await this.store.api.user.me();
    this.save();
  }
}
