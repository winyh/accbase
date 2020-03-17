import { Card, Table } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import React, { Component } from 'react';

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
  },
  {
    title: '用户id',
    dataIndex: 'user_id',
  },
  {
    title: '用户姓名',
    dataIndex: 'name',
    defaultSortOrder: 'name',
  },
  {
    title: '手机',
    dataIndex: 'telephone',
    defaultSortOrder: 'telephone',
  },
  {
    title: '角色',
    dataIndex: 'role',
    defaultSortOrder: 'role',
  },
  {
    title: '操作内容',
    dataIndex: 'content',
    defaultSortOrder: 'content',
  },
  {
    title: '操作类型',
    dataIndex: 'type',
    defaultSortOrder: 'type',
  },
  {
    title: '最后登录ip',
    dataIndex: 'last_ip',
  },
  {
    title: '操作时间',
    dataIndex: 'action_at',
    defaultSortOrder: 'action_at',
  },
];

const data = [
  {
    key: '1',
    id: '1',
    user_id: '23',
    name: '赵子龙',
    telephone: 18672882782,
    role: '普通用户',
    content: '新增用户',
    type: '新增',
    last_ip: '192.278.219.9',
    action_at: '2020-03-08 18:30:20',
  },
  {
    key: '2',
    id: '2',
    user_id: '23',
    name: '刘玄德',
    telephone: 18672882782,
    role: '普通用户',
    content: '新增用户',
    type: '新增',
    last_ip: '192.278.219.9',
    action_at: '2020-03-08 18:30:20',
  },
  {
    key: '3',
    id: '3',
    user_id: '23',
    name: '关云长',
    telephone: 18672882782,
    role: '普通用户',
    content: '新增用户',
    type: '新增',
    last_ip: '192.278.219.9',
    action_at: '2020-03-08 18:30:20',
  },
  {
    key: '4',
    id: '4',
    user_id: '23',
    name: 'winyh',
    telephone: 18672882782,
    role: '普通用户',
    content: '新增用户',
    type: '新增',
    last_ip: '192.278.219.9',
    action_at: '2020-03-08 18:30:20',
  },
];

class LogList extends Component<> {
  onChange = (pagination, filters, sorter, extra) => {
    console.log('params', pagination, filters, sorter, extra);
  };

  render() {
    return (
      <PageHeaderWrapper>
        <Card>
          <Table columns={columns} dataSource={data} onChange={this.onChange} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default LogList;
