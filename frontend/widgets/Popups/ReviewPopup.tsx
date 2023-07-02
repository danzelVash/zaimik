import PopupWrapper from '@/shared/PopupWrapper';
import Button from '@/shared/Button';

const ReviewPopup: React.FC = () => {
	return <PopupWrapper currentPopup='review' className='max-w-[800px]'>
		<div className='relative z-[2]'>
				<div className='text-center lg:text-[50px] md:text-[45px] text-[28px] text-white font-bold leading-tight'>
					Оставить отзыв
				</div>
				<form className='lg:mt-6 md:mt-5 mt-4 lg:space-y-6 md:space-y-5 space-y-4'>
					<div className='w-full lg:h-[75px] md:h-[60px] h-[51px]'>
						<input
							className='w-full h-full bg-primary-light rounded-xl px-6 lg:text-[28px] md:text-[24px] text-[18px] text-dark placeholder:text-[#414141CC]'
							type='text'
							placeholder='Имя'
						/>
					</div>
					<div className='w-full lg:h-[75px] md:h-[60px] h-[51px]'>
						<input
							className='w-full h-full bg-primary-light rounded-xl px-6 lg:text-[28px] md:text-[24px] text-[18px] text-dark placeholder:text-[#414141CC]'
							type='text'
							placeholder='Номер телефона'
						/>
					</div>
					<div className='w-full'>
						<textarea
							className='resize-none md:h-[150px] h-[125px] w-full bg-primary-light rounded-xl px-6 py-4 lg:text-[28px] md:text-[24px] text-[18px] text-dark placeholder:text-[#414141CC]'
							placeholder='Комментарий'
						/>
					</div>
					<Button className='rounded-xl' type='submit'>Отправить</Button>
					<div className='text-center text-[#000000bf] lg;text-[28px] md:text-[24px] text-[16px] leading-tight font-semibold'>
						Ваш номер телефона будет скрыт
					</div>
				</form>
			</div>
	</PopupWrapper>
};

export default ReviewPopup;