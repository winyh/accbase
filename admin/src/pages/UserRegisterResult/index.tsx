import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import { Button, Result } from 'antd';
import Link from 'umi/link';
import React from 'react';
import { RouteChildrenProps } from 'react-router';

import styles from './style.less';

const actions = (
  <div className={styles.actions}>
    <a href="">
      <Button size="large" type="primary">
        <FormattedMessage id="userregisterresult.register-result.view-mailbox" />
      </Button>
    </a>
    <Link to="/">
      <Button size="large">
        <FormattedMessage id="userregisterresult.register-result.back-home" />
      </Button>
    </Link>
  </div>
);

const UserRegisterResult: React.FC<RouteChildrenProps> = ({ location }) => (
  <Result
    className={styles.registerResult}
    status="success"
    title={
      <div className={styles.title}>
        <FormattedMessage
          id="userregisterresult.register-result.msg"
          values={{ email: location.state ? location.state.account : 'AntDesign@example.com' }}
        />
      </div>
    }
    subTitle={formatMessage({ id: 'userregisterresult.register-result.activation-email' })}
    extra={actions}
  />
);

export default UserRegisterResult;
