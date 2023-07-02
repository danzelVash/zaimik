import Button from '@/shared/Button';
import PopupWrapper from '@/shared/PopupWrapper';

const AuthPopup: React.FC = () => {
	return (
		<PopupWrapper currentPopup='auth' className='max-w-[800px]'>
			<div className='relative z-[2]'>
				<div className='text-center lg:text-[50px] md:text-[45px] text-[28px] text-white font-bold leading-tight'>
					Авторизация
				</div>
				<form className='lg:mt-6 md:mt-5 mt-4 lg:space-y-6 md:space-y-5 space-y-4'>
					<div className='w-full lg:h-[75px] md:h-[60px] h-[51px]'>
						<input
							className='w-full h-full bg-primary-light rounded-xl px-6 lg:text-[28px] md:text-[24px] text-[18px] text-dark placeholder:text-[#414141CC]'
							type='text'
							placeholder='Email'
						/>
					</div>
					<div>
						<Button className='rounded-xl'>Получить код</Button>
					</div>
					<div className='w-full lg:h-[75px] md:h-[60px] h-[51px]'>
						<input
							className='w-full h-full bg-primary-light rounded-xl px-6 lg:text-[28px] md:text-[24px] text-[18px] text-dark placeholder:text-[#414141CC]'
							type='text'
							placeholder='Введите код'
						/>
					</div>
					<div>
						<Button className='rounded-xl'>Войти</Button>
					</div>
					<div className='text-center underline text-[#000000b3] lg;text-[28px] md:text-[24px] text-[18px] leading-tight font-semibold'>
						<span className='cursor-pointer hover:opacity-70 transition-opacity duration-200'>
							Прислать код повторно
						</span>
					</div>
				</form>
			</div>
		</PopupWrapper>
	);
};

export default AuthPopup;
