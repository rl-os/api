import { AxiosInstance } from "axios";
import { Token } from "../models/oauth";

type Scope = '*' | 'read.*'
type GrantType = 'password'

interface GrantTypePwd {
  username: string;
  password: string;
}

export class AuthAPI {
  constructor(
    private readonly axios: AxiosInstance,
    private readonly clientId: number,
    private readonly clientSecret: string
  ) {}

  // tslint:disable-next-line:variable-name
  public async token<T extends GrantTypePwd>(scope: Scope, grant_type: GrantType, body: T): Promise<Token> {
    const { data } = await this.axios.post(
      "/oauth/token",
      {
        client_id: this.clientId,
        client_secret: this.clientSecret,
        grant_type,
        scope,
        ...body
      }
    );

    return data;
  }

  public async loginByPwd(username: string, password: string): Promise<Token> {
    return this.token('*', 'password', {
      username,
      password
    });
  }
}
