import Kid from '../../../assets/images/kid.png'
import KidWithBag from '../../../assets/images/kid_with_bag.png'

function Images() {
  return (
    <div className="images absolute hidden md:grid grid-cols-2 gap-4 w-full mt-8">
      <img src={Kid} alt="Kid" className="w-full max-w-2xl h-auto" />
      <img
        src={KidWithBag}
        alt="Kid with Bag"
        className="w-full max-w-2xl h-auto ml-auto"
      />
    </div>
  )
}

export default Images
