import {Avatar, Tooltip} from "@nextui-org/react";
import {Badge} from "@nextui-org/badge";
import {css, cx} from "@emotion/css";
import dayjs from "dayjs";
import {UseAvatarProps} from "@nextui-org/avatar/dist/use-avatar";

interface CustomAvatarProps extends UseAvatarProps {
    tooltip: boolean;
    username?: string;
    email?: string;
    onlineTime?: number;
    colorIndex?: number;
}

export function avatarParseInitials(username: string | undefined): string {
    if (!username) {
        return '';
    }
    let [firstName = '', secondName = ''] = username
        .split(/[^A-Za-zА-Яа-я]/)
        .map((word: string): string => word.replace(/[^A-Za-zА-Яа-я]/, ''))
        .filter(Boolean);
    if (firstName === '') {
        return '';
    }
    if (secondName === '') {
        return firstName[0];
    }
    return firstName[0] + secondName[0];
}

const avatarColors = [
    {
        'name': 'mango',
        'from': '#FFB457',
        'to': '#FF705B',
        'textColor': 'white',
    },
    {
        'name': 'green',
        'from': '#11998E',
        'to': '#38EF7D',
        'textColor': 'white',
    },
    {
        'name': 'Reaqua',
        'from': '#799F0CFF',
        'to': '#ACBB78FF',
        'textColor': 'white',
    },
    {
        'name': 'Bluelagoo',
        'from': '#0052D4',
        'to': '#4364F7',
        'textColor': 'white',
    },
    {
        'name': 'Anwar',
        'from': '#334D50FF',
        'to': '#CBCAA5FF',
        'textColor': 'white',
    },
    {
        'name': 'Blu',
        'from': '#00416AFF',
        'to': '#E4E5E6FF',
        'textColor': 'white',
    },
    {
        'name': 'PiggyPink',
        'from': '#EE9CA7FF',
        'to': '#FFDDE1FF',
        'textColor': 'white',
    },
    {
        'name': 'CoolBlues',
        'from': '#2193B0FF',
        'to': '#6DD5EDFF',
        'textColor': 'white',
    },
    {
        'name': 'MoonlitAsteroid',
        'from': '#0F2027FF',
        'to': '#203A43FF',
        'textColor': 'white',
    },
    {
        'name': 'DarkOcean',
        'from': '#373B44FF',
        'to': '#4286F4FF',
        'textColor': 'white',
    },
    {
        'name': 'Yoda',
        'from': '#FF0099FF',
        'to': '#493240FF',
        'textColor': 'white',
    },
    {
        'name': 'Amin',
        'from': '#8E2DE2FF',
        'to': '#4A00E0FF',
        'textColor': 'white',
    },
    {
        'name': 'Harvey',
        'from': '#1F4037FF',
        'to': '#99F2C8FF',
        'textColor': 'white',
    },
    {
        'name': 'Neuromancer',
        'from': '#F953C6FF',
        'to': '#B91D73FF',
        'textColor': 'white',
    },
    {
        'name': 'Flare',
        'from': '#F12711FF',
        'to': '#F5AF19FF',
        'textColor': 'white',
    },
    {
        'name': 'UltraVoilet',
        'from': '#654EA3FF',
        'to': '#EAAFC8FF',
        'textColor': 'white',
    },
    {
        'name': 'BurningOrange',
        'from': '#FF416CFF',
        'to': '#FF4B2BFF',
        'textColor': 'white',
    },
    {
        'name': 'SinCityRed',
        'from': '#ED213AFF',
        'to': '#93291EFF',
        'textColor': 'white',
    },
    {
        'name': 'BlueRaspberry',
        'from': '#00B4DBFF',
        'to': '#0083B0FF',
        'textColor': 'white',
    },
    {
        'name': 'Vanusa',
        'from': '#DA4453FF',
        'to': '#89216BFF',
        'textColor': 'white',
    },
    {
        'name': 'Vanusa',
        'from': '#AD5389FF',
        'to': '#3C1053FF',
        'textColor': 'white',
    },
    {
        'name': 'MoonPurple',
        'from': '#4E54C8FF',
        'to': '#8F94FBFF',
        'textColor': 'white',
    },
    {
        'name': 'Bighead',
        'from': '#C94B4BFF',
        'to': '#4B134FFF',
        'textColor': 'white',
    },
    {
        'name': 'Selenium',
        'from': '#3C3B3FFF',
        'to': '#605C3CFF',
        'textColor': 'white',
    },
    {
        'name': 'PinkFlavour',
        'from': '#800080FF',
        'to': '#FFC0CBFF',
        'textColor': 'white',
    },
    {
        'name': 'OrangeFun',
        'from': '#FC4A1AFF',
        'to': '#F7B733FF',
        'textColor': 'white',
    },
    {
        'name': 'DigitalWater',
        'from': '#74EBD5FF',
        'to': '#ACB6E5FF',
        'textColor': 'white',
    },
    {
        'name': 'Lithium',
        'from': '#6D6027FF',
        'to': '#D3CBB8FF',
        'textColor': 'white',
    }
];

function getHash(value: string): number {
    return value
        .split('')
        .reduce((a, b) => ((a << 5) - a + b.charCodeAt(0)) | 0, 0);
}

function avatarParseColor(
    username: string | undefined = undefined,
    colorIndex: number | undefined = undefined
): string {
    if (!username) return 'default';
    const colorNumber = colorIndex ?? getHash(username);
    const color = avatarColors[colorNumber % avatarColors.length];
    return css({
        backgroundImage: `linear-gradient(to bottom right, ${color.from}, ${color.to})`,
        color: color.textColor
    });
}

function avatarIsOnline(onlineTime?: number): boolean | undefined {
    if (onlineTime === undefined) return undefined;
    if (!onlineTime) return false;
    const now = dayjs();
    const past = dayjs.unix(onlineTime);
    const diffMinutes = now.diff(past, 'minute');
    if (diffMinutes < 0) {
        return undefined;
    }
    return diffMinutes < 5;
}

function CustomAvatar(props: CustomAvatarProps) {
    if (!props.username) return <></>;
    const isOnline = avatarIsOnline(props.onlineTime);
    const username = props.email || props.username;
    const avatar = (
        <Avatar
            size={props.size}
            classNames={{
                base: cx(avatarParseColor(username, props.colorIndex))
            }}
            name={avatarParseInitials(username)}
            as={props.as}

        />
    );
    if (!props.tooltip) return avatar;
    return (
        <span className='relative'>
        <Tooltip showArrow={true} content={username}>
          <Badge
              isInvisible={isOnline === undefined}
              size={props.size}
              isDot={true}
              color={isOnline ? 'success' : 'danger'}
              content=""
              shape="circle"
          >
            {avatar}
          </Badge>
        </Tooltip>
      </span>
    );
}

export default CustomAvatar;
