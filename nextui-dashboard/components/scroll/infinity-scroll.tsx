'use client';

import React, {useEffect, useRef} from "react";
import {Loading} from "@/components/scroll/loader";

type Props = {
    isLoading?: boolean;
    children: React.ReactNode;
    loadMore?: () => void;
};

function InfiniteScroll(props: Props) {
    const observerElement = useRef<HTMLDivElement | null>(null);
    const {isLoading, children, loadMore} = props;

    useEffect(() => {
        // is element in view?
        function handleIntersection(entries: IntersectionObserverEntry[]) {
            entries.forEach((entry) => {
                if (entry.isIntersecting && !isLoading) {
                    loadMore && loadMore();
                }
            });
        }

        // create observer instance
        const observer = new IntersectionObserver(handleIntersection, {
            root: null,
            rootMargin: "100px",
            threshold: 0,
        });

        if (observerElement.current) {
            observer.observe(observerElement.current);
        }

        // cleanup function
        return () => {
            observer.disconnect();
            observerElement.current && observerElement.current.scroll({top: 0})
        }
    }, [isLoading, loadMore]);

    return (
        <>
            <>{children}</>

            <div ref={observerElement}>
                {isLoading && (
                    <div className="wrapper flex justify-center items-center h-20">
                        <Loading size={12}/>
                    </div>
                )}
            </div>
        </>
    );
}

export default InfiniteScroll;