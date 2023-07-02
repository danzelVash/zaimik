import Raccoon from '@/shared/Raccoon';
import SearchTitle from '@/shared/SearchTitle';
import ChatConditions from './ChatConditions';
import ChatField from './ChatField';

const Chat: React.FC = () => {
	return (
		<>
			<Raccoon
				src='/static/raccoons/welcoming_raccoon.png'
				width={213}
				height={258}
				className='md:w-auto md:h-auto w-[190px] h-[230px] absolute z-[-1] left-1/2 top-0 -translate-x-1/2 -translate-y-1/2'
			/>
			<SearchTitle>Ваши условия</SearchTitle>
			<ChatConditions />
			<ChatField />
		</>
	);
};

export default Chat;
