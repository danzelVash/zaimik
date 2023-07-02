import {
	ICity,
	IFooterLink,
	IGeneral,
	ILoanCondition,
	INavLink,
	IObtainingStage,
	IPaymentMethod,
} from './types';

export const general: IGeneral = {
	email: 'zaimik666@mail.ru',
};

export const cities: ICity[] = [
	{
		id: 1,
		name: 'Москва',
	},
	{
		id: 2,
		name: 'Санкт-Петербург',
	},
	{
		id: 3,
		name: 'Казань',
	},
	{
		id: 4,
		name: 'Екатеринбург',
	},
];

export const navLinks: INavLink[] = [
	{
		id: 1,
		path: '/',
		name: 'Рассчет микрозайма',
		icon: {
			src: '/icons/calculator.svg',
			size: [28, 36],
		},
	},
	{
		id: 2,
		path: '/reviews',
		name: 'Отзывы',
		icon: {
			src: '/icons/reviews.svg',
			size: [46, 27],
		},
	},
];

export const footerLinks: IFooterLink[] = [
	{
		id: 1,
		path: '/privacy-policy',
		name: 'Политика конфиденциальности',
	},
	{
		id: 2,
		path: '/user-agreement',
		name: 'Пользовательское соглашение',
	},
	{
		id: 3,
		path: '',
		name: 'Необходимые условия для займа',
	},
	{
		id: 4,
		path: '',
		name: 'Отзывы',
	},
];

export const loanConditions: ILoanCondition[] = [
	{
		id: 1,
		icon: {
			src: '/static/home/LoanConditionEntity/icon_1.svg',
			size: [135, 186],
		},
		title: 'Возраст',
		text: 'Для одобрения займа необходимо быть старше 18 лет',
	},
	{
		id: 2,
		icon: {
			src: '/static/home/LoanConditionEntity/icon_2.svg',
			size: [155, 153],
		},
		title: 'Работа',
		text: 'Наличие постоянного места работы необязательно',
	},
	{
		id: 3,
		icon: {
			src: '/static/home/LoanConditionEntity/icon_3.svg',
			size: [110, 180],
		},
		title: 'Паспорт',
		text: 'Для получения займов необходимо иметь гражданство РФ',
	},
];

export const paymentMethods: IPaymentMethod[] = [
	{
		id: 1,
		src: '/static/home/PaymentRulesEntity/visa.svg',
		size: [143, 46],
	},
	{
		id: 2,
		src: '/static/home/PaymentRulesEntity/mastercard.svg',
		size: [83, 51],
	},
	{
		id: 3,
		src: '/static/home/PaymentRulesEntity/mir.svg',
		size: [176, 48],
	},
];

export const obtainingStages: IObtainingStage[] = [
	{
		id: 1,
		styles: 'lg:items-start items-center lg:text-left',
		name: 'Заполните анкету',
		icon: {
			src: '/static/home/ObtainigStagesEntity/icon_1.svg',
			size: [106, 106],
		},
	},
	{
		id: 2,
		styles: 'items-center lg:text-center text-left',
		name: 'выберите кредитное предложение',
		icon: {
			src: '/static/home/ObtainigStagesEntity/icon_2.svg',
			size: [106, 106],
		},
	},
	{
		id: 3,
		styles: 'lg:items-end items-center lg:text-right text-left',
		name: 'получите деньги',
		icon: {
			src: '/static/home/ObtainigStagesEntity/icon_3.svg',
			size: [106, 106],
		},
	},
];
