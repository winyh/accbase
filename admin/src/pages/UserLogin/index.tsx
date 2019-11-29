import { Alert, Checkbox, Icon } from 'antd';
import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import React, { Component } from 'react';

import { CheckboxChangeEvent } from 'antd/es/checkbox';
import { Dispatch } from 'redux';
import { FormComponentProps } from 'antd/es/form';
import Link from 'umi/link';
import { connect } from 'dva';
import { StateType } from './model';
import LoginComponents from './components/Login';
import styles from './style.less';

const { Tab, UserName, Password, Mobile, Captcha, Submit } = LoginComponents;

interface UserLoginProps {
  dispatch: Dispatch<any>;
  userLogin: StateType;
  submitting: boolean;
}
interface UserLoginState {
  type: string;
  autoLogin: boolean;
}
export interface FormDataType {
  userName: string;
  password: string;
  mobile: string;
  captcha: string;
}

@connect(
  ({
    userLogin,
    loading,
  }: {
    userLogin: StateType;
    loading: {
      effects: {
        [key: string]: string;
      };
    };
  }) => ({
    userLogin,
    submitting: loading.effects['userLogin/login'],
  }),
)
class UserLogin extends Component<UserLoginProps, UserLoginState> {
  loginForm: FormComponentProps['form'] | undefined | null = undefined;

  state: UserLoginState = {
    type: 'account',
    autoLogin: true,
  };

  changeAutoLogin = (e: CheckboxChangeEvent) => {
    this.setState({
      autoLogin: e.target.checked,
    });
  };

  handleSubmit = (err: any, values: FormDataType) => {
    const { type } = this.state;
    if (!err) {
      const { dispatch } = this.props;
      dispatch({
        type: 'userLogin/login',
        payload: {
          ...values,
          type,
        },
      });
    }
  };

  onTabChange = (type: string) => {
    this.setState({ type });
  };

  onGetCaptcha = () =>
    new Promise((resolve, reject) => {
      if (!this.loginForm) {
        return;
      }
      this.loginForm.validateFields(['mobile'], {}, (err: any, values: FormDataType) => {
        if (err) {
          reject(err);
        } else {
          const { dispatch } = this.props;
          ((dispatch({
            type: 'userLogin/getCaptcha',
            payload: values.mobile,
          }) as unknown) as Promise<any>)
            .then(resolve)
            .catch(reject);
        }
      });
    });

  renderMessage = (content: string) => (
    <Alert style={{ marginBottom: 24 }} message={content} type="error" showIcon />
  );

  render() {
    const { userLogin, submitting } = this.props;
    const { status, type: loginType } = userLogin;
    const { type, autoLogin } = this.state;
    return (
      <div className={styles.main}>
        <LoginComponents
          defaultActiveKey={type}
          onTabChange={this.onTabChange}
          onSubmit={this.handleSubmit}
          ref={(form: any) => {
            this.loginForm = form;
          }}
        >
          <Tab key="account" tab={formatMessage({ id: 'userlogin.login.tab-login-credentials' })}>
            {status === 'error' &&
              loginType === 'account' &&
              !submitting &&
              this.renderMessage(
                formatMessage({ id: 'userlogin.login.message-invalid-credentials' }),
              )}
            <UserName
              name="userName"
              placeholder={`${formatMessage({ id: 'userlogin.login.userName' })}: admin or user`}
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userlogin.userName.required' }),
                },
              ]}
            />
            <Password
              name="password"
              placeholder={`${formatMessage({ id: 'userlogin.login.password' })}: ant.design`}
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userlogin.password.required' }),
                },
              ]}
              onPressEnter={e => {
                e.preventDefault();
                this.loginForm.validateFields(this.handleSubmit);
              }}
            />
          </Tab>
          <Tab key="mobile" tab={formatMessage({ id: 'userlogin.login.tab-login-mobile' })}>
            {status === 'error' &&
              loginType === 'mobile' &&
              !submitting &&
              this.renderMessage(
                formatMessage({ id: 'userlogin.login.message-invalid-verification-code' }),
              )}
            <Mobile
              name="mobile"
              placeholder={formatMessage({ id: 'userlogin.phone-number.placeholder' })}
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userlogin.phone-number.required' }),
                },
                {
                  pattern: /^1\d{10}$/,
                  message: formatMessage({ id: 'userlogin.phone-number.wrong-format' }),
                },
              ]}
            />
            <Captcha
              name="captcha"
              placeholder={formatMessage({ id: 'userlogin.verification-code.placeholder' })}
              countDown={120}
              onGetCaptcha={this.onGetCaptcha}
              getCaptchaButtonText={formatMessage({ id: 'userlogin.form.get-captcha' })}
              getCaptchaSecondText={formatMessage({ id: 'userlogin.captcha.second' })}
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userlogin.verification-code.required' }),
                },
              ]}
            />
          </Tab>
          <div>
            <Checkbox checked={autoLogin} onChange={this.changeAutoLogin}>
              <FormattedMessage id="userlogin.login.remember-me" />
            </Checkbox>
            <a style={{ float: 'right' }} href="">
              <FormattedMessage id="userlogin.login.forgot-password" />
            </a>
          </div>
          <Submit loading={submitting}>
            <FormattedMessage id="userlogin.login.login" />
          </Submit>
          <div className={styles.other}>
            <FormattedMessage id="userlogin.login.sign-in-with" />
            <Icon type="alipay-circle" className={styles.icon} theme="outlined" />
            <Icon type="taobao-circle" className={styles.icon} theme="outlined" />
            <Icon type="weibo-circle" className={styles.icon} theme="outlined" />
            <Link className={styles.register} to="/user/register">
              <FormattedMessage id="userlogin.login.signup" />
            </Link>
          </div>
        </LoginComponents>
      </div>
    );
  }
}

export default UserLogin;
