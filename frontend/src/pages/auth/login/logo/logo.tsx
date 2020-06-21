import React from 'react';
import classNames from "classnames";
import { observer } from 'mobx-react-lite';

import styles from './logo.module.scss';

const Logo = () => {
  return <div className={classNames(styles.root, "d-flex flex-column-auto flex-column")}>
    <a href="/" className="text-center mb-10">
      <img
        src="https://301222.selcdn.ru/akasi/assets/logo/logo.svg"
        className="max-h-150px"
        alt="risu.life"
      />
    </a>
    <h3 className={classNames(styles.text)}>
      The first private server for osu!lazer,<br/>
      created for players by players
    </h3>
    <h3 className={classNames(styles.text)}>
      We have a closed community <br/>
      in which there is no place for cheaters and inadequate people
    </h3>
  </div>;
};

export default observer(Logo);
