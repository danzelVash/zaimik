import Raccoon from '@/shared/Raccoon';
import SearchText from '@/shared/SearchText';
import SearchTitle from '@/shared/SearchTitle';
import Image from 'next/image';
import LoadingCard from './LoadingCard';

const Loading: React.FC = () => {
	return (
		<div>
			<Raccoon
				src='/static/raccoons/thoughtful_raccoon.png'
				width={213}
				height={258}
				className='md:w-auto md:h-auto w-[190px] h-[230px] absolute z-[-1] left-1/2 top-0 -translate-x-1/2 -translate-y-1/2'
			/>
			<SearchTitle>Я обрабатываю вашу заявку на займ</SearchTitle>
			<SearchText className='lg:mt-4 mt-3'>
				Наши алгоритмы обрабатывают введенную вами информацию{' '}
				<br className='lg:block hidden' /> и подготавливают подходящие для вас
				предложения.
			</SearchText>
			<Image
				className='mx-auto lg:mt-6 md:mt-5 mt-4 md:w-auto md:h-auto w-[40px] h-[40px]'
				src='/icons/clock.svg'
				width={60}
				height={60}
				alt=''
			/>
			<LoadingCard />
		</div>
	);
};

export default Loading;
