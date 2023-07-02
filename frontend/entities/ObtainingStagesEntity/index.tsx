import Title from '@/shared/Title';
import ObtainingStagesItem from '@/components/ObtainingStagesItem';

const ObtainingStagesEntity: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md'>
				<Title>
					Как <span className='text-tertiary'>получить</span> займ?
				</Title>
				<ObtainingStagesItem />
			</div>
		</section>
	);
};

export default ObtainingStagesEntity;
