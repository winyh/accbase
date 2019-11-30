import { Card, Table, Button } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import React, { Component } from 'react';

const columns = [
  {
    title: '菜单名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '路由地址',
    dataIndex: 'href',
    key: 'href',
  },
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type',
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action',
  },
];

const data = [
  {
    key: 1,
    name: '欢迎页',
    href: '/welcome',
    type: '顶级目录',
    status: 0,
    action: '删除',
  },
  {
    key: 2,
    name: '管理页',
    href: '/admin',
    type: '顶级目录',
    status: 1,
    action: '删除',
  },
  {
    key: 3,
    name: '用户管理',
    href: '/sms',
    type: '顶级目录',
    status: 0,
    action: '删除',
    children: [
      {
        key: 31,
        name: '用户列表',
        href: '/sms/userlist',
        type: '菜单',
        status: 1,
        action: '删除',
      },
      {
        key: 32,
        name: '角色列表',
        href: '/sms/rolelist',
        type: '菜单',
        status: 1,
        action: '删除',
      },
      {
        key: 33,
        name: '组织机构',
        href: '/sms/company',
        type: '菜单',
        status: 1,
        action: '删除',
      },
    ],
  },
  {
    key: 4,
    name: '权限管理',
    href: '/role',
    type: '顶级目录',
    status: 1,
    action: '删除',
    children: [
      {
        key: 41,
        name: '用户列表',
        href: '/sms/userlist',
        type: '菜单',
        status: 1,
        action: '删除',
      },
      {
        key: 42,
        name: '角色列表',
        href: '/sms/rolelist',
        type: '菜单',
        status: 1,
        action: '删除',
      },
    ],
  },
  {
    key: 5,
    name: '系统管理',
    href: '/system',
    type: '顶级目录',
    status: 1,
    action: '删除',
    children: [
      {
        key: 51,
        name: '系统信息',
        href: '/system/info',
        type: '菜单',
        status: 1,
        action: '删除',
      },
      {
        key: 52,
        name: '日志列表',
        href: '/system/loglist',
        type: '菜单',
        status: 1,
        action: '删除',
      },
    ],
  },
];

// rowSelection objects indicates the need for row selection
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
  onSelect: (record, selected, selectedRows) => {
    console.log(record, selected, selectedRows);
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    console.log(selected, selectedRows, changeRows);
  },
};

class Company extends Component<> {
  render() {
    return (
      <PageHeaderWrapper>
        <Card>
          <Button icon="plus" type="primary" style={{ marginBottom: '20px' }}>
            新建
          </Button>
          <Table columns={columns} rowSelection={rowSelection} dataSource={data} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default Company;
