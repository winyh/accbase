import React from 'react';
import { DefaultFooter } from '@ant-design/pro-layout';

const year = new Date().getFullYear();

export default () => (
  <DefaultFooter
    copyright={`2019~${year} winyh`}
    links={[
      {
        key: 'Accbase',
        title: 'Accbase',
        href: 'https://github.com/winyh/accbase',
        blankTarget: true,
      },
    ]}
  />
);
