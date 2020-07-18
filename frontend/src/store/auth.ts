import { observable } from "mobx";
import { Store } from ".";
import { DetailedUser } from "../models/detailed_user";
import Axios from "axios";
import { KJUR } from 'jsrsasign';

const LOCALSTORAGE_KEY = 'user_auth';

export class Auth {
  @observable
  public currentUser: DetailedUser | null = null;

  @observable
  public accessToken: string | null = null;

  @observable
  public refreshToken: string | null = null;

  private tokenWatcher: TokenWatcher | null = null;

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

    this.tokenWatcher = new TokenWatcher(this.store);

    try {
      await this.tokenWatcher.start();
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
    this.accessToken = null;
    this.refreshToken = null;

    this.save();
  }

  public async login(username: string, pwd: string): Promise<void> {
    const token = await this.store.api.auth.loginByPwd(username, pwd);

    this.accessToken = token.access_token;
    this.refreshToken = token.refresh_token;

    this.currentUser = await this.store.api.user.me();
    this.save();
  }
}

class TokenWatcher {
  private readonly cancelToken = Axios.CancelToken.source();

  constructor(
    private readonly store: Store
  ) {}

  private async onIssue() {}

  public start() {
    return this.safeLoop(() => this.check());
  }

  private async check(): Promise<void> {
    if (this.store.auth.accessToken === null) return;

    console.log(KJUR.jws.JWS.parse(this.store.auth.accessToken).payloadObj);
  }

  private async safeLoop(fn: () => Promise<void>) {
    while (this.cancelToken.token.reason === undefined) {
      try {
        await fn();
        await new Promise(resolve => setTimeout(resolve, 5 * 1000));
      } catch (e) {
        if (!Axios.isCancel(e)) {
          await new Promise(resolve => setTimeout(resolve, 30 * 1000));
        }
      }
    }
  }

}
