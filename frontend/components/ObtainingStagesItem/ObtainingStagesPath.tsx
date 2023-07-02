import { obtainingStages } from '@/constants';
import Image from 'next/image';
import ObtainingStagesLine from './ObtainingStagesLine';

const ObtainingStagesPath: React.FC = () => {
	return (
		<div
			className={`lg:mt-4 grid lg:grid-cols-3 grid-cols-1 lg:grid-rows-1 grid-rows-3 gap-x-4 md:gap-y-[196px] gap-y-[130px] relative lg:w-full`}
		>
			{obtainingStages?.map((obj, index) => {
				const [width, height] = obj.icon.size;
				return (
					<>
						<div
							key={obj.id}
							className={`flex lg:flex-col flex-row-reverse justify-end gap-y-3 md:gap-x-4 gap-x-3 ${obj.styles}`}
						>
							<div className='xl:text-[28px] md:text-[21px] text-[17px] leading-tight font-bold uppercase'>
								{obj.name}
							</div>
							<div className='shrink-0 md:w-[106px] md:h-[106px] w-[50px] h-[50px]'>
								<Image
									className='max-w-full'
									src={obj.icon.src}
									width={width}
									height={height}
									alt=''
								/>
							</div>
						</div>
						{index + 1 < obtainingStages.length && (
							<ObtainingStagesLine index={index} />
						)}
					</>
				);
			})}
		</div>
	);
};

export default ObtainingStagesPath;
