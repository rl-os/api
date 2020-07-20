import React from 'react';
import classNames from "classnames";
import { observer } from 'mobx-react-lite';

import styles from './logo.module.scss';
import { useTranslate } from "../../../../utils/hooks";

const Logo = () => {
  const { t } = useTranslate("auth");

  return <div className={classNames(styles.root, "d-flex flex-column-auto flex-column")}>
    <a href="/" className="text-center mb-10">
      <img
        src="https://301222.selcdn.ru/akasi/assets/rl-logo/logo-white.svg"
        className="max-h-150px"
        alt="risu.life"
      />
    </a>
    <h3 className={classNames(styles.text)}>
      {t`sinopsis_1`}
    </h3>
    <h3 className={classNames(styles.text)}>
      {t`sinopsis_2`}
    </h3>
  </div>;
};

export default observer(Logo);
