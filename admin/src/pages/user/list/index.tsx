import { Card, Table, Button } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import React, { Component } from 'react';

const columns = [
  {
    title: '姓名',
    dataIndex: 'name',
    filters: [
      {
        text: 'Joe',
        value: 'Joe',
      },
      {
        text: 'Jim',
        value: 'Jim',
      },
      {
        text: 'Submenu',
        value: 'Submenu',
        children: [
          {
            text: 'Green',
            value: 'Green',
          },
          {
            text: 'Black',
            value: 'Black',
          },
        ],
      },
    ],
    // specify the condition of filtering result
    // here is that finding the name started with `value`
    onFilter: (value, record) => record.name.indexOf(value) === 0,
    sorter: (a, b) => a.name.length - b.name.length,
    sortDirections: ['descend'],
  },
  {
    title: '手机',
    dataIndex: 'telephone',
    defaultSortOrder: 'telephone',
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    defaultSortOrder: 'email',
  },
  {
    title: '年龄',
    dataIndex: 'age',
    defaultSortOrder: 'descend',
    sorter: (a, b) => a.age - b.age,
  },
  {
    title: '角色',
    dataIndex: 'role',
    defaultSortOrder: 'role',
  },
  {
    title: '组织部门',
    dataIndex: 'company',
    defaultSortOrder: 'company',
  },
  {
    title: '地址',
    dataIndex: 'address',
    filters: [
      {
        text: 'London',
        value: 'London',
      },
      {
        text: 'New York',
        value: 'New York',
      },
    ],
    filterMultiple: false,
    onFilter: (value, record) => record.address.indexOf(value) === 0,
    sorter: (a, b) => a.address.length - b.address.length,
    sortDirections: ['descend', 'ascend'],
  },
];

const data = [
  {
    key: '1',
    name: '赵子龙',
    telephone: 18672882782,
    email: '2712192172@qq.com',
    age: 32,
    role: '普通用户',
    company: '武汉分公司',
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '2',
    name: '刘玄德',
    telephone: 18672882782,
    email: '2712192172@qq.com',
    age: 42,
    role: '管理员',
    company: '合肥分公司',
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '3',
    name: '关云长',
    telephone: 18672882782,
    email: '2712192172@qq.com',
    age: 32,
    role: '管理员',
    company: '黄石分公司',
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '4',
    name: 'winyh',
    telephone: 18672882782,
    email: '2712192172@qq.com',
    age: 32,
    role: '管理员',
    company: '黄石分公司',
    address: 'New York No. 1 Lake Park',
  },
];

class UserList extends Component<> {
  onChange = (pagination, filters, sorter, extra) => {
    console.log('params', pagination, filters, sorter, extra);
  };

  render() {
    return (
      <PageHeaderWrapper>
        <Card>
          <Button icon="plus" type="primary" style={{ marginBottom: '20px' }}>
            新建
          </Button>
          <Table columns={columns} dataSource={data} onChange={this.onChange} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default UserList;
