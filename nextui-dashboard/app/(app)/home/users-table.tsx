import {Link as LinkComponent} from "@nextui-org/link";
import {Link} from "react-router-dom";
import {UsersTableWrapper} from "@/components/pages/accounts/table/table";
import React from "react";
import {useUsersList} from "@/helpers/queries/users/get";

const HomeUsersWidget = () => {
    const response = useUsersList({limit: 5});
    const users = response?.data?.pages.flatMap(
        (p) => p.result?.model ?? []
    ) || [];
    return (
        <div className="flex flex-col justify-center w-full py-5 px-4 lg:px-0  max-w-[90rem] mx-auto gap-3">
            <div className="flex  flex-wrap justify-between">
                <h3 className="text-center text-xl font-semibold">Latest Users</h3>
                <LinkComponent href="#" color="primary"
                               className="cursor-pointer">
                    <Link
                        to="/accounts"
                    >
                        View All
                    </Link>
                </LinkComponent>
            </div>
            <UsersTableWrapper
                users={users}
                isLoading={response.isLoading}
            />
        </div>
    );
};

export default HomeUsersWidget;
