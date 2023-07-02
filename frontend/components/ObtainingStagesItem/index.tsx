import Raccoon from '@/shared/Raccoon';
import ObtainingStagesPath from './ObtainingStagesPath';

const ObtainingStagesItem: React.FC = () => {
	return (
		<div className='lg:mt-4 md:mt-7 mt-4 flex lg:flex-col flex-row-reverse items-center justify-between'>
			<Raccoon
				src='/static/raccoons/thoughtful_raccoon.png'
				width={273}
				height={328}
				className='shrink-0 md:w-auto md:h-auto w-[131px] h-[157px]'
			/>
			<ObtainingStagesPath />
		</div>
	);
};

export default ObtainingStagesItem;
