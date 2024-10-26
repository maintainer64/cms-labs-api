import clsx from "clsx";
import {Providers} from "@/app/providers";


export const App = () => {
    return (
        <div className={clsx("font-sans antialiased")}>
            <Providers/>
        </div>
    )
}
