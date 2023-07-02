'use client';

import Button from '@/shared/Button';
import ReviewCard from '@/components/ReviewCard';
import Image from 'next/image';
import { PopupContext } from '@/context/PopupContext';
import { useContext } from 'react';

const ReviewEntity: React.FC = () => {
	const ctx = useContext(PopupContext);

	return (
		<div>
			<div className='grid lg:grid-cols-2 grid-cols-1 lg:gap-8 md:gap-6 gap-4 lg:mt-6 md:mt-5 mt-4'>
				{Array(6)
					.fill('')
					.map((item, index) => (
						<ReviewCard key={index} />
					))}
			</div>
			<div className='lg:mt-6 md:mt-5 mt-4'>
				<div
					onClick={() => ctx.setPopup('review')}
					className='lg:max-w-[600px] md:max-w-[344px] mx-auto'>
					<Button>
						Оставить отзыв
					</Button>
				</div>
				<button
					className='md:w-auto w-full flex flex-row items-center justify-center mx-auto gap-x-2.5 lg:mt-6 mt-4 text-[#000000b3] bg-accent rounded-3xl md:py-4 py-3 md:px-6 px-3 font-medium lg:text-[21px] md:text-[20px] text-[18px] transition-colors duration-300 hover:bg-[#dc7c2a]'
				>
					<span>Загрузить больше отзывов</span>
					<Image className='md:w-[18px] w-[16px]' src="/icons/down.svg" width={18} height={11} alt='' />
				</button>
			</div>
		</div>
	);
};

export default ReviewEntity;
