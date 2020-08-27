import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Alert, Typography } from 'antd';
import styles from './Welcome.less';

const CodePreview: React.FC<{}> = ({ children }) => (
  <pre className={styles.pre}>
    <code>
      <Typography.Text copyable>{children}</Typography.Text>
    </code>
  </pre>
);

export default (): React.ReactNode => (
  <PageContainer>
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
        权限管理{' '}
        <a href="https://protable.ant.design/" rel="noopener noreferrer" target="__blank">
          基于 go-micro 开发的微服务
        </a>
      </Typography.Text>
      <CodePreview>git clone https://github.com/winyh/accbase</CodePreview>
    </Card>
  </PageContainer>
);
