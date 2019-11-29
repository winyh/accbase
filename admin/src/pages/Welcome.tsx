import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { FormattedMessage } from 'umi-plugin-react/locale';
import { Card, Typography, Alert } from 'antd';

import styles from './Welcome.less';

const CodePreview: React.FC<{}> = ({ children }) => (
  <pre className={styles.pre}>
    <code>
      <Typography.Text copyable>{children}</Typography.Text>
    </code>
  </pre>
);

export default (): React.ReactNode => (
  <PageHeaderWrapper>
    <Card>
      <Alert
        message="Accbase 是基于 Go-micro Casbin Gorm JWT 开发，以微服务的方式构建权限管理系统"
        type="success"
        showIcon
        banner
        style={{
          margin: -12,
          marginBottom: 24,
        }}
      />
      <Typography.Text strong>
        <a target="_blank" rel="noopener noreferrer" href="https://github.com/winyh/accbase">
          <FormattedMessage
            id="app.welcome.link.block-list"
            defaultMessage="基于 Go-micro Casbin Gorm 开发，以微服务的方式构建权限管理系统"
          />
        </a>
      </Typography.Text>
      <CodePreview>git clone https://github.com/winyh/accbase</CodePreview>
    </Card>
  </PageHeaderWrapper>
);
