import { Card, Descriptions, Badge, Typography } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import React, { Component } from 'react';

const CodePreview: React.FC<{}> = ({ children }) => (
  <pre>
    <code>
      <Typography.Text copyable>{children}</Typography.Text>
    </code>
  </pre>
);

class SystemInfo extends Component<> {
  render() {
    return (
      <PageHeaderWrapper>
        <Card>
          <Descriptions title="系统信息" bordered>
            <Descriptions.Item label="项目名称">Accbase 权限管理</Descriptions.Item>
            <Descriptions.Item label="开发者">winyh</Descriptions.Item>
            <Descriptions.Item label="邮箱">2712192471@qq.com</Descriptions.Item>
            <Descriptions.Item label="创建时间">2019-08-24 18:00:00</Descriptions.Item>
            <Descriptions.Item label="Usage Time" span={2}>
              2019-04-24 18:00:00
            </Descriptions.Item>
            <Descriptions.Item label="当前状态" span={3}>
              <Badge status="processing" text="开发迭代中" />
            </Descriptions.Item>
            <Descriptions.Item label="本项目开发语言">Golang TypeScript HTML5</Descriptions.Item>
            <Descriptions.Item label="项目成员">1</Descriptions.Item>
            <Descriptions.Item label="是否开源">是</Descriptions.Item>
            <Descriptions.Item label="配置信息">
              Go version: v1.4
              <br />
              go micro: v1.18
              <br />
              UI 界面: Antd-Design-Pro v4.0
              <br />
              技术栈: Casbin + Gorm + Gin + Go + TypeScript
              <br />
              项目地址: <CodePreview>https://github.com/winyh/accbase</CodePreview>
              <br />
            </Descriptions.Item>
          </Descriptions>
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default SystemInfo;
