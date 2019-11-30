import { Card, Table, Button } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import React, { Component } from 'react';

const columns = [
  {
    title: '角色名称',
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
    title: '描述',
    dataIndex: 'descritpion',
    defaultSortOrder: 'descritpion',
    sorter: (a, b) => a.age - b.age,
  },
  {
    title: '状态',
    dataIndex: 'age',
    defaultSortOrder: 'descend',
    sorter: (a, b) => a.age - b.age,
  },
  {
    title: 'Address',
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
    name: '超级管理员',
    description: '平台，集团，总公司，连锁机构',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '2',
    name: '管理员',
    description: '分公司，区域',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '3',
    name: '财务人员',
    description: '也有所属公司和关系',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '4',
    name: '运行人员',
    description: '平台，集团，总公司',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '5',
    name: '编辑',
    description: '平台信息发布，基础信息录入',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '6',
    name: '门店',
    description: '公司机构下属门店',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '7',
    name: '员工',
    description: '阿姨月嫂',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '8',
    name: '客户',
    description: '平台客户',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '9',
    name: '客服',
    description: '客服，售后处理人员',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '10',
    name: '销售',
    description: '产品销售',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
];

class RoleList extends Component<> {
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

export default RoleList;
