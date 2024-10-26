import {useAlert} from "./hooks";
import Alert from "@/components/alerts/alert";

interface GroupAlertsProps {
    children: React.ReactNode
}

export const GroupAlerts = ({children}: GroupAlertsProps) => {
    const {hideAlert, alertMessages} = useAlert();
    return (
        <div>
            <div className="fixed top-12 right-4 space-y-4 z-50 w-1/4">
                {alertMessages.map((alert, index) => (
                    <Alert
                        key={index}
                        type={alert.type}
                        message={alert.message}
                        description={alert.description}
                        onClose={hideAlert.bind(hideAlert, index)}
                    />
                ))}
            </div>
            <div>{children}</div>
        </div>
    );
};

export default GroupAlerts;
