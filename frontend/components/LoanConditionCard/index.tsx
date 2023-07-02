import { ILoanCondition } from '@/constants/types';
import Image from 'next/image';

const LoanConditionCard: React.FC<ILoanCondition> = ({
	id,
	icon,
	title,
	text,
}) => {
	const [imageWidth, imageHeight] = icon.size;

	return (
		<div
			className={`xl:col-span-2 ${
				id % 3 !== 0 ? 'md:col-span-3' : ''
			} col-span-6 main-gradient flex lg:flex-row flex-col gap-5 items-center lg:p-6 p-5 rounded-3xl text-white`}
		>
			<div
				className={`w-[${imageWidth}px] h-[${imageHeight}px] md:min-h-[186px] flex items-center justify-center shrink-0`}
			>
				<Image
					className='w-full h-full'
					src={icon.src}
					width={imageWidth}
					height={imageHeight}
					alt=''
				/>
			</div>
			<div className='lg:self-auto self-start'>
				<div className='font-extrabold lg:text-[40px] text-[36px] leading-tight'>
					{title}
				</div>
				<div className='lg:mt-3 mt-2 font-medium lg:text-[20px] text-[18px] leading-tight'>
					{text}
				</div>
			</div>
		</div>
	);
};

export default LoanConditionCard;
