import {PanelHeader, PanelHeaderBack} from "@vkontakte/vkui";
import {AdminPanelMenu, AdminPanelMenuTitle} from "@/views/admin/base.tsx";
import {useDataProvider} from "@colday/shared/src/components/viewParams/hooks.ts";
import VkAdminFormReleaseEdit from "@colday/shared/src/components/adminPanel/formReleaseEdit";
import {ReleasesSchemaItem} from "@colday/api";
import {editRelease} from "@colday/shared/src/queries/siteInfo/editRelease";
import {usePhotosInfo} from "@colday/shared/src/queries/siteInfo/getPhotosInfo";
import {useAlert} from "@colday/shared/src/components/alerts/hooks";


type VkAdminAboutPlayAudioProps = {
    id: AdminPanelMenu
    to: (panel: AdminPanelMenu) => void
}
const VkAdminReleaseEdit = ({id, to}: VkAdminAboutPlayAudioProps) => {
    const photos = usePhotosInfo();
    const {data} = useDataProvider<ReleasesSchemaItem>();
    const {hideAlert, showAlert} = useAlert();
    const {mutate} = editRelease({
        onMutate: () => {
            hideAlert()
        },
        onSuccess: (data, editData) => {
            if (!editData.id) {
                to(AdminPanelMenu.releases)
                return
            }
            showAlert({
                type: "default",
                message: data.status,
            })
        },
        onError: () => {
            showAlert({
                type: "error",
                message: "Произошла ошибка при обновлении сайта"
            })
        }
    })
    return <>
        <PanelHeader before={
            <PanelHeaderBack label="Назад" onClick={to.bind(to, AdminPanelMenu.releases)}/>
        }>{AdminPanelMenuTitle.get(id)}</PanelHeader>
        <VkAdminFormReleaseEdit
            values={data}
            photos={photos.data}
            isLoading={photos.isLoading}
            onSubmit={mutate}
            onDelete={() => {

            }}
        />
    </>
}
export default VkAdminReleaseEdit;
