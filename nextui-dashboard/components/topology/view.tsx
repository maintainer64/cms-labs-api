import React from 'react';
import {Background, Controls, ReactFlow} from '@xyflow/react';
import {edgeTypes, nodeTypes} from './objectTypes';
import {getLayoutElements} from './autoLayout';
import {useTopologyGet} from "@/helpers/queries/topology/get";
import {useParams} from "react-router-dom";
import {RFEdgeTopology, RFNodeTopology} from "@/components/topology/objectTypes/types";
import {Loading} from "@/components/scroll/loader";
import {ErrorModal} from "@/components/pages/auth/error";
import {Button} from "@heroui/react";

interface Props {
    children: React.ReactNode;
}

export const Layout = ({children}: Props) => {
    return <div style={{height: 'calc(100vh - 70px)', width: '100%'}}>{children}</div>;
};

export const TopologyFlowVisualization = () => {
    const {namespace} = useParams();
    const queryTopology = useTopologyGet(namespace);
    const initNodes = (queryTopology.data?.result?.topology?.nodes ?? []) as RFNodeTopology[];
    const initEdges = (queryTopology.data?.result?.topology?.edges ?? []) as RFEdgeTopology[];
    const {nodes, edges} = getLayoutElements(
        initNodes,
        initEdges,
        "TB"
    );
    console.log({nodes, edges})
    if (queryTopology.isLoading) return <Loading size='md'/>;
    if (queryTopology.error) {
        // @ts-ignore
        const errMsg = queryTopology?.error?.body?.msg || "Внутрянняя ошибка";
        return (
            <ErrorModal title={"Отображение топологии"}
                        description={errMsg}>
                <Button onPress={() => queryTopology.refetch()} href='#' variant='light' color='primary'>
                    Попробовать снова
                </Button>
            </ErrorModal>
        );
    }
    return (
        <Layout>
            <ReactFlow
                nodes={nodes}
                edges={edges}
                nodeTypes={nodeTypes}
                edgeTypes={edgeTypes}
                fitView
            >
                <Background/>
                <Controls/>
            </ReactFlow>
        </Layout>
    );
};
