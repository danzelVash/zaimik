import Raccoon from '@/shared/Raccoon';

const Greeting = () => {
	return (
		<div
			className='
			 lg:mt-6 md:mt-5 mt-3 flex flex-row items-center md:gap-2 gap-1 relative
		'
		>
			<Raccoon
				src='/static/raccoons/welcoming_raccoon.png'
				width={253}
				height={307}
				className='shrink-0 md:w-auto md:h-auto w-[130px] h-[158px]'
			/>
			<div className='bg-primary-light font-bold xl:text-[34px] lg:text-[28px] md:text-[26px] text-[16px] leading-tight uppercase lg:p-5 md:p-3 p-2 rounded-3xl md:border-[5px] border-[3px] border-black'>
				<p className='md:block hidden'>
					Привет, меня зовут <span className='text-primary'>Займик</span>! Я
					помогу подобрать тебе самые выгодные условия займов от разных
					компаний. Оставь заявку ниже и мы предложим тебе подходящий вариант в
					кратчайшие сроки!
				</p>
				<p className='md:hidden block'>
					Привет, меня зовут <span className='text-primary'>Займик</span>! Я
					помогу подобрать тебе самые выгодные условия займов от разных
					компаний.
				</p>
			</div>
		</div>
	);
};

export default Greeting;
