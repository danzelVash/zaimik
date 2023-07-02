import SearchWrapper from '@/shared/SearchWrapper';
import Chat from '@/components/Chat';

export const metadata = {
	title: 'Чат',
};

const ChatPage: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md container-full md:pt-[129px] pt-[115px]'>
				<SearchWrapper>
					<Chat />
				</SearchWrapper>
			</div>
		</section>
	);
};

export default ChatPage;
