import "./StatisticCard.scss"

const StatisticCard = ({ title, value } : { title: string,  value: string}) => {
    return (
        <div className="card">
            <p className="card-title">{title}</p>
            <p className="card-value">{value}</p>
        </div>
    )
}

export default StatisticCard