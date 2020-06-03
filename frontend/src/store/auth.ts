import { observable } from "mobx";
import { Store } from ".";
import { DetailedUser } from "../models/detailed_user";

export class Auth {
  @observable
  public currentUser: DetailedUser | null = null;

  @observable
  public accessToken: string | null = null;

  @observable
  public refreshToken: string | null = null;

  constructor(private readonly store: Store) {}

  public reset() {
    this.currentUser = null;
  }

  public async login(username: string, pwd: string): Promise<void> {
    const token = await this.store.api.auth.loginByPwd(username, pwd);

    this.accessToken = token.accessToken;
    this.refreshToken = token.refreshToken;
  }
}
