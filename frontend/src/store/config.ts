import { observable } from "mobx";
import { Store } from ".";


export class Config {
  @observable
  public language: string = "ru";

  @observable
  public baseAPI: string = "https://osu.ppy.sh"

  @observable
  public clientId: number = 5;

  @observable
  public clientSecret: string = "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk";

  constructor(private readonly store: Store) {}

  public reset() {
    this.language = "ru";
  }
}
