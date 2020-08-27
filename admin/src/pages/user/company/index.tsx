import { Card, Table, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import React, { Component } from 'react';

<PlusOutlined />;

const columns = [
  {
    title: '组织机构',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '员工人数',
    dataIndex: 'age',
    key: 'age',
    width: '12%',
  },
  {
    title: '地址',
    dataIndex: 'address',
    width: '30%',
    key: 'address',
  },
];

const data = [
  {
    key: 1,
    name: '湖北信胜合保有限公司',
    age: 60,
    address: '湖北省武汉市江汉区',
    children: [
      {
        key: 11,
        name: '武汉分公司',
        age: 42,
        address: '湖北省武汉市东湖高新区',
        children: [
          {
            key: 121,
            name: '黄石分公司',
            age: 16,
            address: '湖北省黄石市阳新区',
          },
        ],
      },
      {
        key: 12,
        name: '合肥分公司',
        age: 30,
        address: '安徽省合肥市',
        children: [
          {
            key: 121,
            name: '研发部门',
            age: 16,
            address: '安徽省合肥市技术研发中心',
          },
          {
            key: 121,
            name: '市场部门',
            age: 16,
            address: '安徽省合肥市市场推广部门',
          },
          {
            key: 121,
            name: '运维部门',
            age: 16,
            address: '安徽省合肥市运维部门',
          },
          {
            key: 121,
            name: '财务部门',
            age: 16,
            address: '安徽省合肥市财务部门',
          },
          {
            key: 121,
            name: '总裁办',
            age: 16,
            address: '安徽省合肥市总裁办',
          },
        ],
      },
      {
        key: 13,
        name: '北京分公司',
        age: 72,
        address: '北京市朝阳新区',
        children: [
          {
            key: 131,
            name: '研发部门',
            age: 42,
            address: 'London No. 2 Lake Park',
            children: [
              {
                key: 1311,
                name: 'CMS研发小组',
                age: 25,
                address: 'London No. 3 Lake Park',
              },
              {
                key: 1312,
                name: 'AI人工智能研发小组',
                age: 18,
                address: 'London No. 4 Lake Park',
              },
            ],
          },
        ],
      },
    ],
  },
  {
    key: 2,
    name: '湖北中弘发展集团公司',
    age: 32,
    address: '湖北省武汉市洪山区',
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
          <Button icon={<PlusOutlined />} type="primary" style={{ marginBottom: '20px' }}>
            新建
          </Button>
          <Table columns={columns} rowSelection={rowSelection} dataSource={data} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default Company;
