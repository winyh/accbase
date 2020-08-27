import { Card, Table, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import React, { Component } from 'react';

const columns = [
  {
    title: '权限名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '请求地址',
    dataIndex: 'href',
    key: 'href',
  },
  {
    title: '请求类型',
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
    name: '用户相关',
    href: '/api/user',
    type: 'ANY',
    status: 0,
    action: '删除',
    children: [
      {
        key: 11,
        name: '用户新增',
        href: '/api/user',
        type: 'POST',
        status: 1,
        action: '删除',
      },
      {
        key: 12,
        name: '用户删除',
        href: '/api/user',
        type: 'DELETE',
        status: 1,
        action: '删除',
      },
      {
        key: 13,
        name: '用户修改',
        href: '/api/user',
        type: 'PUT',
        status: 1,
        action: '删除',
      },
      {
        key: 14,
        name: '用户查询',
        href: '/api/user',
        type: 'GET',
        status: 1,
        action: '删除',
      },
    ],
  },
  {
    key: 2,
    name: '菜单相关',
    href: '/api/menu',
    type: 'ANY',
    status: 0,
    action: '删除',
    children: [
      {
        key: 21,
        name: '菜单新增',
        href: '/api/user',
        type: 'POST',
        status: 1,
        action: '删除',
      },
      {
        key: 22,
        name: '菜单删除',
        href: '/api/user',
        type: 'DELETE',
        status: 1,
        action: '删除',
      },
      {
        key: 23,
        name: '菜单修改',
        href: '/api/user',
        type: 'PUT',
        status: 1,
        action: '删除',
      },
      {
        key: 24,
        name: '菜单查询',
        href: '/api/user',
        type: 'GET',
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

class PowerList extends Component<> {
  render() {
    return (
      <PageHeaderWrapper>
        <Card>
          <Button icon={<PlusOutlined />} type="primary" style={{ marginBottom: '20px' }}>
            新建
          </Button>
          <Table columns={columns} rowSelection={rowSelection} dataSource={data} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default PowerList;
