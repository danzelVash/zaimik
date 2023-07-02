import Link from 'next/link';
import ChatFieldInputs from './ChatFieldInputs';
import ChatFieldMessage from './ChatFieldMessage';

const ChatField: React.FC = () => {
	return (
		<div className='mt-7'>
			<div className='scrollbar scrollbar-thumb-tertiary scrollbar-track-[#D9D9D9] lg:p-10 md:p-6 p-4 md:space-y-5 space-y-3 lg:h-[800px] md:h-[750px] h-[395px] overflow-y-auto rounded-3xl shadow-[inset_6px_6px_40px_10px_rgba(0,0,0,.25)] bg-[#41414133]'>
				<ChatFieldMessage className='bg-primary-light'>
					Привет, меня зовут <span className='text-primary'>Займик</span>! Я
					помогу подобрать тебе самые выгодные условия займов от разных
					компаний.
				</ChatFieldMessage>
				<ChatFieldMessage className='bg-primary-light'>
					Напиши свое имя, место работы, сумму займа, срок выплаты, и я подберу
					тебе подходящие варианты
				</ChatFieldMessage>
				<ChatFieldMessage className='bg-white' position='end'>
					<p>Иванов Иван Иванович</p>
					<p>Инженер-технолог</p>
					<p>Займ на сумму 9999 рублей</p>
				</ChatFieldMessage>
				<ChatFieldMessage className='bg-primary-light'>
					Отлично! Я принял ваши данные и начал их обрабатывать. Теперь
					необходимо принять условия соглашения для получения средств и ответить
					на несколько вопросов
				</ChatFieldMessage>
				<ChatFieldMessage className='bg-primary-light'>
					Я ознакомлен и согласен{' '}
					<Link
						className='text-tertiary'
						target='_blank'
						href='/privacy-policy'
					>
						со следующим
					</Link>
					.
				</ChatFieldMessage>
				<ChatFieldMessage className='bg-primary-light'>
					Даю согласие{' '}
					<Link
						className='text-tertiary'
						target='_blank'
						href='/user-agreement'
					>
						хранение учетных данных владельца банковской карты
					</Link>
					.
				</ChatFieldMessage>
			</div>
			<div className='lg:mt-7 md:mt-6 mt-5'>
				<ChatFieldInputs />
				<div className='text-center lg:mt-6 md:mt-5 mt-4'>
					<Link
						className='lg:text-[28px] md:text-[24px] text-[18px] underline transition-opacity duration-200 hover:opacity-70'
						href='/'
					>
						Вернуть на главную
					</Link>
				</div>
			</div>
		</div>
	);
};

export default ChatField;
