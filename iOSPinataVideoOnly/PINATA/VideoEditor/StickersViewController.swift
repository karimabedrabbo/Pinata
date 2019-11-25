//
//  StickersViewController.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/18/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//

import UIKit

class StickersViewController: UIViewController, UIGestureRecognizerDelegate {

    
    let screenSize = UIScreen.main.bounds.size
    
    var headerView: UIView = {
        let header = UIView()
        header.translatesAutoresizingMaskIntoConstraints = false
        return header
    }()
    var holdView: UIView = {
        let hold = UIView()
        hold.translatesAutoresizingMaskIntoConstraints = false
        hold.layer.cornerRadius = 3

        return hold
    }()
    
    var scrollView: UIScrollView = {
        let scroll = UIScrollView()
        scroll.translatesAutoresizingMaskIntoConstraints = false
        scroll.isPagingEnabled = true
        scroll.showsHorizontalScrollIndicator = true
        scroll.contentMode = .center
        return scroll
    }()
    var pageControl: UIPageControl = {
        let page = UIPageControl()
        page.translatesAutoresizingMaskIntoConstraints = false
        page.numberOfPages = 2
        
        return page
    }()
    
    var emojisDelegate: EmojisCollectionViewDelegate = {
        let emojisDel = EmojisCollectionViewDelegate()
        return emojisDel
    }()
    
    var emojisCollectionView: UICollectionView = {
        let emojislayout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
        emojislayout.sectionInset = UIEdgeInsets(top: 20, left: 10, bottom: 10, right: 10)
        emojislayout.itemSize = CGSize(width: 70, height: 70)
        let emojisCollectionView = UICollectionView(frame: .zero, collectionViewLayout: emojislayout)
//        emojisCollectionView.collectionViewLayout = emojislayout
        emojisCollectionView.backgroundColor = .clear
        emojisCollectionView.translatesAutoresizingMaskIntoConstraints = false
        return emojisCollectionView
    }()
    
    var collectionView: UICollectionView = {
        let screenSize = UIScreen.main.bounds.size
        let layout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
        layout.sectionInset = UIEdgeInsets(top: 20, left: 10, bottom: 10, right: 10)
        let width = (CGFloat) ((screenSize.width - 30) / 3.0)
        layout.itemSize = CGSize(width: width, height: 100)
        let collectionView = UICollectionView(frame: .zero, collectionViewLayout: layout)
//        collectionView.collectionViewLayout = layout
        collectionView.backgroundColor = .clear
        collectionView.translatesAutoresizingMaskIntoConstraints = false
        return collectionView
    }()
    
    
    
    var stickers : [UIImage] = []
    weak var stickersViewControllerDelegate : StickersViewControllerDelegate?


    override func viewDidLoad() {
        super.viewDidLoad()

       
        
        collectionView.delegate = self
        collectionView.dataSource = self
        
        
        collectionView.register(StickerCollectionViewCell.self, forCellWithReuseIdentifier: "StickerCollectionViewCell")
        
        //-----------------------------------
        
       
        
        
        emojisCollectionView.delegate = emojisDelegate
        emojisCollectionView.dataSource = emojisDelegate
        
        emojisCollectionView.register(EmojiCollectionViewCell.self, forCellWithReuseIdentifier: "EmojiCollectionViewCell")
        
       
        
        self.view.addSubview(self.headerView)
        
        self.view.addSubview(self.scrollView)
        
        self.scrollView.delegate = self
        
        self.headerView.addSubview(self.holdView)
        self.headerView.addSubview(self.pageControl)
        
       
        
        self.scrollView.addSubview(collectionView)
        
        self.scrollView.addSubview(emojisCollectionView)
        
        NSLayoutConstraint.activate([
            self.headerView.topAnchor.constraint(equalTo: self.view.topAnchor),
            self.headerView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
            self.headerView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
            self.headerView.heightAnchor.constraint(lessThanOrEqualToConstant: 40),
            
            self.holdView.centerXAnchor.constraint(equalTo: self.headerView.centerXAnchor),
            self.holdView.topAnchor.constraint(equalTo: self.headerView.topAnchor, constant: 8),
            self.holdView.widthAnchor.constraint(equalToConstant: 50),
            self.holdView.heightAnchor.constraint(equalToConstant: 5),
            
            self.pageControl.topAnchor.constraint(equalTo: self.holdView.bottomAnchor),
            
            self.pageControl.centerXAnchor.constraint(equalTo: self.headerView.centerXAnchor),
            
            self.scrollView.topAnchor.constraint(equalTo: self.headerView.bottomAnchor),
            self.scrollView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
            self.scrollView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
            self.scrollView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
            
            //self.collectionView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
            self.collectionView.leadingAnchor.constraint(equalTo: self.scrollView.leadingAnchor),
            self.collectionView.widthAnchor.constraint(equalTo: self.scrollView.widthAnchor),
            self.collectionView.heightAnchor.constraint(equalTo: self.scrollView.heightAnchor),
            self.collectionView.topAnchor.constraint(equalTo: self.scrollView.topAnchor),
            
            
            self.emojisCollectionView.leadingAnchor.constraint(equalTo: self.collectionView.trailingAnchor),
            self.emojisCollectionView.widthAnchor.constraint(equalTo: self.scrollView.widthAnchor),
            self.emojisCollectionView.heightAnchor.constraint(equalTo: self.scrollView.heightAnchor),
            self.emojisCollectionView.topAnchor.constraint(equalTo: self.scrollView.topAnchor),
            ])
        
        
        

    }
    
 
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        prepareBackgroundView()
    }
    

    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
    

        scrollView.contentSize = CGSize(width: 2.0 * screenSize.width,
                                        height: scrollView.frame.size.height)


    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    

    
    func prepareBackgroundView(){
        let blurEffect = UIBlurEffect.init(style: .light)
        let visualEffect = UIVisualEffectView.init(effect: blurEffect)
        let bluredView = UIVisualEffectView.init(effect: blurEffect)
        bluredView.contentView.addSubview(visualEffect)
        visualEffect.frame = UIScreen.main.bounds
        bluredView.frame = UIScreen.main.bounds
        view.insertSubview(bluredView, at: 0)
    }
    
    func setStickersViewControllerDelegate(stickersDelegate : StickersViewControllerDelegate) {
        self.stickersViewControllerDelegate = stickersDelegate
        emojisDelegate.stickersViewControllerDelegate = stickersViewControllerDelegate
    }
    
}

extension StickersViewController: UIScrollViewDelegate {
    
    func scrollViewDidScroll(_ scrollView: UIScrollView) {
        let pageWidth = scrollView.bounds.width
        let pageFraction = scrollView.contentOffset.x / pageWidth
        self.pageControl.currentPage = Int(round(pageFraction))
    }
    
    func gestureRecognizer(gestureRecognizer: UIGestureRecognizer, shouldRecognizeSimultaneouslyWithGestureRecognizer otherGestureRecognizer: UIGestureRecognizer) -> Bool {
        return false
    }
}

// MARK: - UICollectionViewDataSource
extension StickersViewController: UICollectionViewDataSource, UICollectionViewDelegate, UICollectionViewDelegateFlowLayout {
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return stickers.count
    }
    
    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        stickersViewControllerDelegate?.didSelectImage(image: stickers[indexPath.item])
    }
    
    func numberOfSections(in collectionView: UICollectionView) -> Int {
        return 1
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let identifier = "StickerCollectionViewCell"
        let cell  = collectionView.dequeueReusableCell(withReuseIdentifier: identifier, for: indexPath) as! StickerCollectionViewCell
        cell.stickerImage.image = stickers[indexPath.item]
        return cell
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, minimumLineSpacingForSectionAt section: Int) -> CGFloat {
        return 4
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, minimumInteritemSpacingForSectionAt section: Int) -> CGFloat {
        return 0
    }
    
    
}
