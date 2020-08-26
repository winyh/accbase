import { Card, Table, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import React, { Component } from 'react';

const columns = [
  {
    title: '角色名称',
    dataIndex: 'name',
  },
  {
    title: '描述',
    dataIndex: 'description',
    defaultSortOrder: 'description',
  },
  {
    title: '状态',
    dataIndex: 'age',
    defaultSortOrder: 'descend',
  },
  {
    title: 'Address',
    dataIndex: 'address',
  },
];

const data = [
  {
    key: '1',
    name: '超级管理员',
    description: '平台，集团，总公司，连锁机构',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '2',
    name: '管理员',
    description: '分公司，区域',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '3',
    name: '财务人员',
    description: '也有所属公司和关系',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '4',
    name: '运行人员',
    description: '平台，集团，总公司',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '5',
    name: '编辑',
    description: '平台信息发布，基础信息录入',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '6',
    name: '门店',
    description: '公司机构下属门店',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '7',
    name: '员工',
    description: '阿姨月嫂',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '8',
    name: '客户',
    description: '平台客户',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '9',
    name: '客服',
    description: '客服，售后处理人员',
    age: 32,
    address: '安徽省合肥市长港区',
  },
  {
    key: '10',
    name: '销售',
    description: '产品销售',
    age: 32,
    address: '安徽省合肥市长港区',
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
          <Button icon={<PlusOutlined />} type="primary" style={{ marginBottom: '20px' }}>
            新建
          </Button>
          <Table columns={columns} dataSource={data} onChange={this.onChange} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default RoleList;
