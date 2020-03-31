import React, { FC, useCallback, useEffect, useMemo, useState } from 'react';
import { Link, RouteComponentProps } from 'react-router-dom';
import { Collapse } from 'antd';
import { ColumnProps } from 'antd/lib/table';

import ClusterCard from './Cluster';
import ITable from '@src/components/ITable';

import { useCapture } from '@src/hooks';
import { formatDate } from '@src/tools';
import { fetchApp, fetchClusters } from '@src/services';
import { renderDeleteWithLinkButton } from '@src/renders';
import { AppQuery, App as AppType, Cluster, defaultBaseModel } from '@src/typings';

export interface AppProps extends RouteComponentProps<{ appName: string; clusterName?: string }> {}

const App: FC<AppProps> = ({ match }) => {
  const [collapseKey, setCollapseKey] = useState<string>('clusters');
  const [appName, setAppName] = useState<string>('');
  const [clusterName, setClusterName] = useState<string>();

  // TODO: appState
  const [, getApps] = useCapture<AppType, AppQuery>({
    fn: fetchApp,
    initialState: { ...defaultBaseModel, appName: '' },
  });
  const [clustersState, getClusters] = useCapture<Cluster[], AppQuery>({ fn: fetchClusters, initialState: [] });

  useEffect(() => {
    const { appName, clusterName } = match.params;
    setAppName(appName);
    if (clusterName) {
      setClusterName(clusterName);
      setCollapseKey('');
    }
  }, [match.params]);

  useEffect(() => {
    if (appName) {
      getApps({ appName });
      getClusters({ appName });
    }
  }, [appName, getApps, getClusters]);

  const columns = useMemo(() => {
    const columns: ColumnProps<Cluster>[] = [
      {
        title: '集群名',
        key: 'clusterName',
        dataIndex: 'clusterName',
        render: clusterName => <Link to={`/apps/${appName}/${clusterName}`}>{clusterName}</Link>,
      },
      { title: '描述', key: 'description', dataIndex: 'description' },
      { title: '创建日期', key: 'createdAt', dataIndex: 'createdAt', width: 200, render: formatDate },
      {
        title: '操作',
        key: 'control',
        width: 100,
        render: (_, cluster) => {
          return (
            <div>{renderDeleteWithLinkButton({ label: '删除', popLabel: '确认删除集群', onDelete: () => {} })}</div>
          );
        },
      },
    ];
    return columns;
  }, [appName]);

  const onFilterKey = useCallback((key: string, item: Cluster) => {
    return item.clusterName.includes(key);
  }, []);

  return (
    <div>
      <Collapse
        className="containerLayout"
        accordion
        activeKey={collapseKey}
        onChange={setCollapseKey as (key: string | string[]) => void}
      >
        <Collapse.Panel key="clusters" header="集群列表">
          <ITable<Cluster>
            rowKey="id"
            bordered={false}
            loading={clustersState.loading}
            dataSource={clustersState.data}
            columns={columns}
            pagination={false}
            showSearch={{ onFilter: onFilterKey }}
          />
        </Collapse.Panel>
      </Collapse>
      {clusterName && <ClusterCard appName={appName} clusterName={clusterName} />}
    </div>
  );
};

export default App;
