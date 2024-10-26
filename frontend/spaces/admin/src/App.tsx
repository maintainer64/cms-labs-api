import AdminPage from "@/views/admin/admin.tsx";
import {VkAdminViewProvider} from "@colday/shared/src/components/viewParams/context";
import {AlertProvider} from "@colday/shared/src/components/alerts/context";


export const App = () => {
    return <AlertProvider>
        <VkAdminViewProvider>
            <AdminPage/>
        </VkAdminViewProvider>
    </AlertProvider>
}
