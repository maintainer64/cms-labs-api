import {PanelHeader, PanelHeaderBack} from "@vkontakte/vkui";
import {AdminPanelMenu, AdminPanelMenuTitle} from "@/views/admin/base.tsx";
import {useGetReleasesList} from "@colday/shared/src/queries/siteInfo/getReleases";
import VkAdminReleasesList from "@colday/shared/src/components/adminPanel/formReleasesList";


type VkAdminAboutPlayAudioProps = {
    id: AdminPanelMenu
    to: (panel: AdminPanelMenu) => void
}
const VkAdminReleases = ({id, to}: VkAdminAboutPlayAudioProps) => {
    const {data, isLoading} = useGetReleasesList(true);
    return <>
        <PanelHeader before={
            <PanelHeaderBack label="Назад" onClick={to.bind(to, AdminPanelMenu.menu)}/>
        }>{AdminPanelMenuTitle.get(id)}</PanelHeader>
        <VkAdminReleasesList
            isLoading={isLoading}
            values={data}
            to={to}
        />
    </>
}
export default VkAdminReleases;
